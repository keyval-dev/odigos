package gateway

import (
	"context"
	"fmt"
	odigosv1 "github.com/keyval-dev/odigos/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	containerName    = "gateway"
	containerImage   = "otel/opentelemetry-collector-contrib:0.55.0"
	containerCommand = "/otelcol"
	confDir          = "/conf"
)

func syncDeployment(dests *odigosv1.DestinationList, gateway *odigosv1.CollectorsGroup, ctx context.Context, c client.Client, scheme *runtime.Scheme) (*appsv1.Deployment, error) {
	logger := log.FromContext(ctx)
	desiredDeployment, err := getDesiredDeployment(dests, gateway, scheme)
	if err != nil {
		logger.Error(err, "Failed to get desired deployment")
		return nil, err
	}

	existing := &appsv1.Deployment{}
	if err := c.Get(ctx, client.ObjectKey{Name: gateway.Name, Namespace: gateway.Namespace}, existing); err != nil {
		if apierrors.IsNotFound(err) {
			logger.V(0).Info("Creating deployment")
			newDeployment, err := createDeployment(desiredDeployment, ctx, c)
			if err != nil {
				logger.Error(err, "failed to create deployment")
				return nil, err
			}
			return newDeployment, nil
		} else {
			logger.Error(err, "failed to get deployment")
			return nil, err
		}
	}

	logger.V(0).Info("Patching deployment")
	newDep, err := patchDeployment(existing, desiredDeployment, ctx, c)
	if err != nil {
		logger.Error(err, "failed to patch deployment")
		return nil, err
	}

	return newDep, nil
}

func createDeployment(desired *appsv1.Deployment, ctx context.Context, c client.Client) (*appsv1.Deployment, error) {
	if err := c.Create(ctx, desired); err != nil {
		return nil, err
	}
	return desired, nil
}

func patchDeployment(existing *appsv1.Deployment, desired *appsv1.Deployment, ctx context.Context, c client.Client) (*appsv1.Deployment, error) {
	updated := existing.DeepCopy()
	if updated.Annotations == nil {
		updated.Annotations = map[string]string{}
	}
	if updated.Labels == nil {
		updated.Labels = map[string]string{}
	}

	updated.Spec = desired.Spec
	updated.ObjectMeta.OwnerReferences = desired.ObjectMeta.OwnerReferences
	for k, v := range desired.ObjectMeta.Annotations {
		updated.ObjectMeta.Annotations[k] = v
	}
	for k, v := range desired.ObjectMeta.Labels {
		updated.ObjectMeta.Labels[k] = v
	}

	patch := client.MergeFrom(existing)
	if err := c.Patch(ctx, updated, patch); err != nil {
		return nil, err
	}

	return updated, nil
}

func getDesiredDeployment(dests *odigosv1.DestinationList, gateway *odigosv1.CollectorsGroup, scheme *runtime.Scheme) (*appsv1.Deployment, error) {
	desiredDeployment := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:      gateway.Name,
			Namespace: gateway.Namespace,
			Labels:    commonLabels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: intPtr(1),
			Selector: &v1.LabelSelector{
				MatchLabels: commonLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: commonLabels,
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{
						{
							Name: configKey,
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gateway.Name,
									},
									Items: []corev1.KeyToPath{
										{
											Key:  configKey,
											Path: fmt.Sprintf("%s.yaml", configKey),
										},
									},
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:    containerName,
							Image:   containerImage,
							Command: []string{containerCommand, fmt.Sprintf("--config=%s/%s.yaml", confDir, configKey)},
							EnvFrom: getSecretsFromDests(dests),
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      configKey,
									MountPath: confDir,
								},
							},
						},
					},
				},
			},
		},
	}

	err := ctrl.SetControllerReference(gateway, desiredDeployment, scheme)
	if err != nil {
		return nil, err
	}

	return desiredDeployment, nil
}

func getSecretsFromDests(destList *odigosv1.DestinationList) []corev1.EnvFromSource {
	var result []corev1.EnvFromSource
	for _, dst := range destList.Items {
		result = append(result, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: dst.Spec.SecretRef.Name,
				},
			},
		})
	}

	return result
}

func intPtr(n int32) *int32 {
	return &n
}
