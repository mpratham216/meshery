package resolver

import (
	"context"

	"github.com/layer5io/meshery/internal/graphql/model"
	"github.com/layer5io/meshery/models"
	"github.com/layer5io/meshery/models/pattern/core"
	meshkitKube "github.com/layer5io/meshkit/utils/kubernetes"
	"github.com/layer5io/meshkit/utils/kubernetes/describe"
	meshsyncmodel "github.com/layer5io/meshsync/pkg/model"
)

func (r *Resolver) getAvailableNamespaces(ctx context.Context, provider models.Provider) ([]*model.NameSpace, error) {
	resourceobjects := make([]meshsyncmodel.ResourceObjectMeta, 0)

	result := provider.GetGenericPersister().Distinct("namespace").Not("namespace = ?", "").Find(&resourceobjects)
	if result.Error != nil {
		r.Log.Error(ErrGettingNamespace(result.Error))
		return nil, result.Error
	}
	namespaces := make([]*model.NameSpace, 0)
	for _, obj := range resourceobjects {
		namespaces = append(namespaces, &model.NameSpace{
			Namespace: obj.Namespace,
		})
	}

	return namespaces, nil
}

// getWorkloads return workloads
func (r *Resolver) getWorkloads(ctx context.Context, name, id *string, trim *bool) (res []*model.OAMCapability, err error) {
	if name != nil && *name != "" && id != nil && *id != "" {
		workload := core.GetWorkloadByID(*name, *id)
		// If trim is set to true then remove the schema from the response
		if trim != nil && *trim {
			workload.OAMRefSchema = ""
		}
		res = append(res, &model.OAMCapability{
			OamDefinition: workload.OAMDefinition,
			OamRefSchema:  &workload.OAMRefSchema,
			Host:          &workload.Host,
			ID:            &workload.ID,
			Restricted:    &workload.Restricted,
			Metadata:      workload.Metadata,
		})
	} else if name != nil && *name != "" {
		workloads := core.GetWorkload(*name)
		for _, w := range workloads {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				w.OAMRefSchema = ""
			}
			res = append(res, &model.OAMCapability{
				OamDefinition: w.OAMDefinition,
				OamRefSchema:  &w.OAMRefSchema,
				Host:          &w.Host,
				ID:            &w.ID,
				Restricted:    &w.Restricted,
				Metadata:      w.Metadata,
			})
		}
	} else {
		// return all the workloads
		workloads := core.GetWorkloads()
		for _, w := range workloads {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				w.OAMRefSchema = ""
			}
			res = append(res, &model.OAMCapability{
				OamDefinition: w.OAMDefinition,
				OamRefSchema:  &w.OAMRefSchema,
				Host:          &w.Host,
				ID:            &w.ID,
				Restricted:    &w.Restricted,
				Metadata:      w.Metadata,
			})
		}
	}
	return
}

// getTraits return traits
func (r *Resolver) getTraits(ctx context.Context, name, id *string, trim *bool) (res []*model.OAMCapability, err error) {
	if name != nil && *name != "" && id != nil && *id != "" {
		trait := core.GetTraitByID(*name, *id)
		// If trim is set to true then remove the schema from the response
		if trim != nil && *trim {
			trait.OAMRefSchema = ""
		}
		res = append(res, &model.OAMCapability{
			OamDefinition: trait.OAMDefinition,
			OamRefSchema:  &trait.OAMRefSchema,
			Host:          &trait.Host,
			ID:            &trait.ID,
			Restricted:    &trait.Restricted,
			Metadata:      trait.Metadata,
		})
	} else if name != nil && *name != "" {
		traits := core.GetWorkload(*name)
		for _, t := range traits {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				t.OAMRefSchema = ""
			}
			res = append(res, &model.OAMCapability{
				OamDefinition: t.OAMDefinition,
				OamRefSchema:  &t.OAMRefSchema,
				Host:          &t.Host,
				ID:            &t.ID,
				Restricted:    &t.Restricted,
				Metadata:      t.Metadata,
			})
		}
	} else {
		// return all the traits
		traits := core.GetWorkloads()
		for _, t := range traits {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				t.OAMRefSchema = ""
			}
			res = append(res, &model.OAMCapability{
				OamDefinition: t.OAMDefinition,
				OamRefSchema:  &t.OAMRefSchema,
				Host:          &t.Host,
				ID:            &t.ID,
				Restricted:    &t.Restricted,
				Metadata:      t.Metadata,
			})
		}
	}

	return
}

// getScopes return scopes
func (r *Resolver) getScopes(ctx context.Context, name, id *string, trim *bool) (res []*model.OAMCapability, err error) {
	if name != nil && *name != "" && id != nil && *id != "" {
		scope := core.GetScopeByID(*name, *id)
		// If trim is set to true then remove the schema from the response
		if trim != nil && *trim {
			scope.OAMRefSchema = ""
		}
		res = append(res, &model.OAMCapability{
			OamDefinition: scope.OAMDefinition,
			OamRefSchema:  &scope.OAMRefSchema,
			Host:          &scope.Host,
			ID:            &scope.ID,
			Restricted:    &scope.Restricted,
			Metadata:      scope.Metadata,
		})
	} else if name != nil && *name != "" {
		scopes := core.GetWorkload(*name)
		for _, s := range scopes {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				s.OAMRefSchema = ""
			}

			res = append(res, &model.OAMCapability{
				OamDefinition: s.OAMDefinition,
				OamRefSchema:  &s.OAMRefSchema,
				Host:          &s.Host,
				ID:            &s.ID,
				Restricted:    &s.Restricted,
				Metadata:      s.Metadata,
			})
		}
	} else {
		// return all the scopes
		scopes := core.GetWorkloads()
		for _, s := range scopes {
			// If trim is set to true then remove the schema from the response
			if trim != nil && *trim {
				s.OAMRefSchema = ""
			}
			res = append(res, &model.OAMCapability{
				OamDefinition: s.OAMDefinition,
				OamRefSchema:  &s.OAMRefSchema,
				Host:          &s.Host,
				ID:            &s.ID,
				Restricted:    &s.Restricted,
				Metadata:      s.Metadata,
			})
		}
	}
	return
}

func (r *Resolver) getKubectlDescribe(ctx context.Context, name string, typeArg string, namespace string) (*model.KctlDescribeDetails, error) {
	var ResourceMap = map[string]describe.DescribeType{
		"Pod":                       describe.Pod,
		"Deployment":                describe.Deployment,
		"Job":                       describe.Job,
		"CronJob":                   describe.CronJob,
		"StatefulSet":               describe.StatefulSet,
		"DaemonSet":                 describe.DaemonSet,
		"ReplicaSet":                describe.ReplicaSet,
		"Secret":                    describe.Secret,
		"Service":                   describe.Service,
		"ServiceAccount":            describe.ServiceAccount,
		"Node":                      describe.Node,
		"LimitRange":                describe.LimitRange,
		"ResourceQuota":             describe.ResourceQuota,
		"PersistentVolume":          describe.PersistentVolume,
		"PersistentVolumeClaim":     describe.PersistentVolumeClaim,
		"Namespace":                 describe.Namespace,
		"Endpoints":                 describe.Endpoints,
		"ConfigMap":                 describe.ConfigMap,
		"PriorityClass":             describe.PriorityClass,
		"Ingress":                   describe.Ingress,
		"Role":                      describe.Role,
		"ClusterRole":               describe.ClusterRole,
		"RoleBinding":               describe.RoleBinding,
		"ClusterRoleBinding":        describe.ClusterRoleBinding,
		"NetworkPolicy":             describe.NetworkPolicy,
		"ReplicationController":     describe.ReplicationController,
		"CertificateSigningRequest": describe.CertificateSigningRequest,
		"EndpointSlice":             describe.EndpointSlice,
	}

	options := describe.DescriberOptions{
		Name:      name,
		Namespace: namespace,
		Type:      ResourceMap[typeArg],
	}

	client, err := meshkitKube.New([]byte(""))
	if err != nil {
		r.Log.Error(ErrMesheryClient(err))
		return nil, err
	}

	details, err := describe.Describe(client, options)
	if err != nil {
		r.Log.Error(ErrKubectlDescribe(err))
		return nil, err
	}

	return &model.KctlDescribeDetails{
		Describe: &details,
	}, nil
}
