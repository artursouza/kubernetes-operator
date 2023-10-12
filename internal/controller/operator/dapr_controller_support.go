package operator

import (
	"context"
	"fmt"
	"strconv"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	ctrlCli "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/dapr-sandbox/dapr-kubernetes-operator/pkg/controller/predicates"
)

func gcSelector(rc *ReconciliationRequest) (labels.Selector, error) {

	namespace, err := labels.NewRequirement(
		DaprReleaseNamespace,
		selection.Equals,
		[]string{rc.Resource.Namespace})

	if err != nil {
		return nil, fmt.Errorf("cannot determine release namespace requirement: %w", err)
	}

	name, err := labels.NewRequirement(
		DaprReleaseName,
		selection.Equals,
		[]string{rc.Resource.Name})

	if err != nil {
		return nil, fmt.Errorf("cannot determine release name requirement: %w", err)
	}

	generation, err := labels.NewRequirement(
		DaprReleaseGeneration,
		selection.LessThan,
		[]string{strconv.FormatInt(rc.Resource.Generation, 10)})

	if err != nil {
		return nil, fmt.Errorf("cannot determine generation requirement: %w", err)
	}

	selector := labels.NewSelector().
		Add(*namespace).
		Add(*name).
		Add(*generation)

	return selector, nil
}

func labelsToRequest(_ context.Context, object ctrlCli.Object) []reconcile.Request {
	allLabels := object.GetLabels()
	if allLabels == nil {
		return nil
	}
	name := allLabels[DaprReleaseName]
	if name == "" {
		return nil
	}
	namespace := allLabels[DaprReleaseNamespace]
	if namespace == "" {
		return nil
	}

	return []reconcile.Request{{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}}
}

func dependantWithLabels(watchUpdate bool, watchDelete bool, watchStatus bool) predicate.Predicate {
	return predicate.And(
		&predicates.HasLabel{
			Name: DaprReleaseName,
		},
		&predicates.HasLabel{
			Name: DaprReleaseNamespace,
		},
		&predicates.DependentPredicate{
			WatchUpdate: watchUpdate,
			WatchDelete: watchDelete,
			WatchStatus: watchStatus,
		},
	)
}

func ReleaseSelector() (labels.Selector, error) {
	hasReleaseNameLabel, err := labels.NewRequirement(DaprReleaseName, selection.Exists, []string{})
	if err != nil {
		return nil, err
	}

	hasReleaseNamespaceLabel, err := labels.NewRequirement(DaprReleaseNamespace, selection.Exists, []string{})
	if err != nil {
		return nil, err
	}

	selector := labels.NewSelector().
		Add(*hasReleaseNameLabel).
		Add(*hasReleaseNamespaceLabel)

	return selector, nil
}

func CurrentReleaseSelector(rc *ReconciliationRequest) (labels.Selector, error) {
	namespace, err := labels.NewRequirement(
		DaprReleaseNamespace,
		selection.Equals,
		[]string{rc.Resource.Namespace})

	if err != nil {
		return nil, fmt.Errorf("cannot determine release namespace requirement: %w", err)
	}

	name, err := labels.NewRequirement(
		DaprReleaseName,
		selection.Equals,
		[]string{rc.Resource.Name})

	if err != nil {
		return nil, fmt.Errorf("cannot determine release name requirement: %w", err)
	}

	generation, err := labels.NewRequirement(
		DaprReleaseGeneration,
		selection.Equals,
		[]string{strconv.FormatInt(rc.Resource.Generation, 10)})

	if err != nil {
		return nil, fmt.Errorf("cannot determine generation requirement: %w", err)
	}

	selector := labels.NewSelector().
		Add(*namespace).
		Add(*name).
		Add(*generation)

	return selector, nil
}
