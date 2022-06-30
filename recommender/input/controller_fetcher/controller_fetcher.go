package controller_fetcher

import (
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/cache"
)

type wellKnownController string

// ControllerKey identifies a controller.
type ControllerKey struct {
	Namespace string
	Kind      string
	Name      string
}

// ControllerKeyWithAPIVersion identifies a controller and API it's defined in.
type ControllerKeyWithAPIVersion struct {
	ControllerKey
	ApiVersion string
}

// ControllerFetcher is responsible for finding the topmost well-known or scalable controller
type ControllerFetcher interface {
	// FindTopMostWellKnownOrScalable returns topmost well-known or scalable controller. Error is returned if controller cannot be found.
	FindTopMostWellKnownOrScalable(controller *ControllerKeyWithAPIVersion) (*ControllerKeyWithAPIVersion, error)
}

type controllerFetcher struct {
	scaleNamespacer              scale.ScalesGetter
	mapper                       apimeta.RESTMapper
	informersMap                 map[wellKnownController]cache.SharedIndexInformer
	scaleSubresourceCacheStorage controllerCacheStorage
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
Only two methods are exported from this file, they are called from cluster_feeder.
NewControllerFetcher
FindTopMostWellKnownOrScalable

Going into details seems not necessary
*/

func (f *controllerFetcher) isWellKnown(key *ControllerKeyWithAPIVersion) bool {
	// check if f.informersMap contains key.Controller.Kind
	return false
}

func (f *controllerFetcher) isWellKnownOrScalable(key *ControllerKeyWithAPIVersion) bool {
	// if inWellKnown() , return true
	// Get the RestMappings for this groupKind from f.mapper
	// f.getScaleForResources()
	return false
}

func (f *controllerFetcher) getOwnerForScaleResource(groupKind schema.GroupKind, namespace, name string) (*ControllerKeyWithAPIVersion, error) {
	return nil, nil
}

func (f *controllerFetcher) FindTopMostWellKnownOrScalable(key *ControllerKeyWithAPIVersion) (*ControllerKeyWithAPIVersion, error) {
	return nil, nil
}
