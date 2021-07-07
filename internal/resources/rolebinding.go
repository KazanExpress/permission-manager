package resources

import (
	"context"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleBindingService interface {
	CreateRole(namespace string, ctx context.Context, rb *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error)
	DeleteRoleBinding(namespace, roleBindingName string, ctx context.Context) error
}

func (r *resourceService) CreateRole(namespace string, ctx context.Context, rb *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {

	rb, err := r.kubeclient.RbacV1().RoleBindings(namespace).Create(ctx, rb, metav1.CreateOptions{})

	if err != nil {
		return nil, err
	}

	return rb, nil

}

func (r *resourceService) DeleteRoleBinding(namespace, roleBindingName string, ctx context.Context) error {

	return r.kubeclient.RbacV1().RoleBindings(namespace).Delete(ctx, roleBindingName, metav1.DeleteOptions{})

}
