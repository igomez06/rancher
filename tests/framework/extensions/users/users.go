package users

import (
	"github.com/rancher/rancher/tests/framework/clients/rancher"
	management "github.com/rancher/rancher/tests/framework/clients/rancher/generated/management/v3"
)

func CreateUserWithRole(rancherClient *rancher.Client, user *management.User, role *management.GlobalRoleBinding) (*management.User, error) {
	createdUser, err := rancherClient.Management.User.Create(user)
	if err != nil {
		return nil, err
	}

	role.UserID = createdUser.ID

	_, err = rancherClient.Management.GlobalRoleBinding.Create(role)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func AddProjectMember(rancherClient *rancher.Client, project *management.Project, user *management.User, projectRole string) (*management.ProjectRoleTemplateBinding, error) {
	role := &management.ProjectRoleTemplateBinding{
		ProjectID:       project.ID,
		UserPrincipalID: user.PrincipalIDs[0],
		RoleTemplateID:  projectRole,
	}

	return rancherClient.Management.ProjectRoleTemplateBinding.Create(role)
}

func RemoveProjectMember(rancherClient *rancher.Client, role *management.ProjectRoleTemplateBinding) error {
	return rancherClient.Management.ProjectRoleTemplateBinding.Delete(role)
}
