package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_Action = map[string]string{
	"":                   "Action describes a request to the API server",
	"namespace":          "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces",
	"verb":               "Verb is one of: get, list, watch, create, update, delete",
	"resourceAPIGroup":   "Group is the API group of the resource Serialized as resourceAPIGroup to avoid confusion with the 'groups' field when inlined",
	"resourceAPIVersion": "Version is the API version of the resource Serialized as resourceAPIVersion to avoid confusion with TypeMeta.apiVersion and ObjectMeta.resourceVersion when inlined",
	"resource":           "Resource is one of the existing resource types",
	"resourceName":       "ResourceName is the name of the resource being requested for a \"get\" or deleted for a \"delete\"",
	"content":            "Content is the actual content of the request for create and update",
}

func (Action) SwaggerDoc() map[string]string {
	return map_Action
}

var map_ClusterPolicy = map[string]string{
	"":             "ClusterPolicy is a object that holds all the ClusterRoles for a particular namespace.  There is at most one ClusterPolicy document per namespace.",
	"metadata":     "Standard object's metadata.",
	"lastModified": "LastModified is the last time that any part of the ClusterPolicy was created, updated, or deleted",
	"roles":        "Roles holds all the ClusterRoles held by this ClusterPolicy, mapped by ClusterRole.Name",
}

func (ClusterPolicy) SwaggerDoc() map[string]string {
	return map_ClusterPolicy
}

var map_ClusterPolicyBinding = map[string]string{
	"":             "ClusterPolicyBinding is a object that holds all the ClusterRoleBindings for a particular namespace.  There is one ClusterPolicyBinding document per referenced ClusterPolicy namespace",
	"metadata":     "Standard object's metadata.",
	"lastModified": "LastModified is the last time that any part of the ClusterPolicyBinding was created, updated, or deleted",
	"policyRef":    "PolicyRef is a reference to the ClusterPolicy that contains all the ClusterRoles that this ClusterPolicyBinding's RoleBindings may reference",
	"roleBindings": "RoleBindings holds all the ClusterRoleBindings held by this ClusterPolicyBinding, mapped by ClusterRoleBinding.Name",
}

func (ClusterPolicyBinding) SwaggerDoc() map[string]string {
	return map_ClusterPolicyBinding
}

var map_ClusterPolicyBindingList = map[string]string{
	"":         "ClusterPolicyBindingList is a collection of ClusterPolicyBindings",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterPolicyBindings",
}

func (ClusterPolicyBindingList) SwaggerDoc() map[string]string {
	return map_ClusterPolicyBindingList
}

var map_ClusterPolicyList = map[string]string{
	"":         "ClusterPolicyList is a collection of ClusterPolicies",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterPolicies",
}

func (ClusterPolicyList) SwaggerDoc() map[string]string {
	return map_ClusterPolicyList
}

var map_ClusterRole = map[string]string{
	"":         "ClusterRole is a logical grouping of PolicyRules that can be referenced as a unit by ClusterRoleBindings.",
	"metadata": "Standard object's metadata.",
	"rules":    "Rules holds all the PolicyRules for this ClusterRole",
}

func (ClusterRole) SwaggerDoc() map[string]string {
	return map_ClusterRole
}

var map_ClusterRoleBinding = map[string]string{
	"":           "ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference any ClusterRole in the same namespace or in the global namespace. It adds who information via (Users and Groups) OR Subjects and namespace information by which namespace it exists in. ClusterRoleBindings in a given namespace only have effect in that namespace (excepting the master namespace which has power in all namespaces).",
	"metadata":   "Standard object's metadata.",
	"userNames":  "UserNames holds all the usernames directly bound to the role. This field should only be specified when supporting legacy clients and servers. See Subjects for further details.",
	"groupNames": "GroupNames holds all the groups directly bound to the role. This field should only be specified when supporting legacy clients and servers. See Subjects for further details.",
	"subjects":   "Subjects hold object references to authorize with this rule. This field is ignored if UserNames or GroupNames are specified to support legacy clients and servers. Thus newer clients that do not need to support backwards compatibility should send only fully qualified Subjects and should omit the UserNames and GroupNames fields. Clients that need to support backwards compatibility can use this field to build the UserNames and GroupNames.",
	"roleRef":    "RoleRef can only reference the current namespace and the global namespace. If the ClusterRoleRef cannot be resolved, the Authorizer must return an error. Since Policy is a singleton, this is sufficient knowledge to locate a role.",
}

func (ClusterRoleBinding) SwaggerDoc() map[string]string {
	return map_ClusterRoleBinding
}

var map_ClusterRoleBindingList = map[string]string{
	"":         "ClusterRoleBindingList is a collection of ClusterRoleBindings",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterRoleBindings",
}

func (ClusterRoleBindingList) SwaggerDoc() map[string]string {
	return map_ClusterRoleBindingList
}

var map_ClusterRoleList = map[string]string{
	"":         "ClusterRoleList is a collection of ClusterRoles",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterRoles",
}

func (ClusterRoleList) SwaggerDoc() map[string]string {
	return map_ClusterRoleList
}

var map_GroupRestriction = map[string]string{
	"":       "GroupRestriction matches a group either by a string match on the group name or a label selector applied to group labels.",
	"groups": "Groups is a list of groups used to match against an individual user's groups. If the user is a member of one of the whitelisted groups, the user is allowed to be bound to a role.",
	"labels": "Selectors specifies a list of label selectors over group labels.",
}

func (GroupRestriction) SwaggerDoc() map[string]string {
	return map_GroupRestriction
}

var map_IsPersonalSubjectAccessReview = map[string]string{
	"": "IsPersonalSubjectAccessReview is a marker for PolicyRule.AttributeRestrictions that denotes that subjectaccessreviews on self should be allowed",
}

func (IsPersonalSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_IsPersonalSubjectAccessReview
}

var map_LocalResourceAccessReview = map[string]string{
	"": "LocalResourceAccessReview is a means to request a list of which users and groups are authorized to perform the action specified by spec in a particular namespace",
}

func (LocalResourceAccessReview) SwaggerDoc() map[string]string {
	return map_LocalResourceAccessReview
}

var map_LocalSubjectAccessReview = map[string]string{
	"":       "LocalSubjectAccessReview is an object for requesting information about whether a user or group can perform an action in a particular namespace",
	"user":   "User is optional.  If both User and Groups are empty, the current authenticated user is used.",
	"groups": "Groups is optional.  Groups is the list of groups to which the User belongs.",
	"scopes": "Scopes to use for the evaluation.  Empty means \"use the unscoped (full) permissions of the user/groups\". Nil for a self-SAR, means \"use the scopes on this request\". Nil for a regular SAR, means the same as empty.",
}

func (LocalSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_LocalSubjectAccessReview
}

var map_NamedClusterRole = map[string]string{
	"":     "NamedClusterRole relates a name with a cluster role",
	"name": "Name is the name of the cluster role",
	"role": "Role is the cluster role being named",
}

func (NamedClusterRole) SwaggerDoc() map[string]string {
	return map_NamedClusterRole
}

var map_NamedClusterRoleBinding = map[string]string{
	"":            "NamedClusterRoleBinding relates a name with a cluster role binding",
	"name":        "Name is the name of the cluster role binding",
	"roleBinding": "RoleBinding is the cluster role binding being named",
}

func (NamedClusterRoleBinding) SwaggerDoc() map[string]string {
	return map_NamedClusterRoleBinding
}

var map_NamedRole = map[string]string{
	"":     "NamedRole relates a Role with a name",
	"name": "Name is the name of the role",
	"role": "Role is the role being named",
}

func (NamedRole) SwaggerDoc() map[string]string {
	return map_NamedRole
}

var map_NamedRoleBinding = map[string]string{
	"":            "NamedRoleBinding relates a role binding with a name",
	"name":        "Name is the name of the role binding",
	"roleBinding": "RoleBinding is the role binding being named",
}

func (NamedRoleBinding) SwaggerDoc() map[string]string {
	return map_NamedRoleBinding
}

var map_Policy = map[string]string{
	"":             "Policy is a object that holds all the Roles for a particular namespace.  There is at most one Policy document per namespace.",
	"metadata":     "Standard object's metadata.",
	"lastModified": "LastModified is the last time that any part of the Policy was created, updated, or deleted",
	"roles":        "Roles holds all the Roles held by this Policy, mapped by Role.Name",
}

func (Policy) SwaggerDoc() map[string]string {
	return map_Policy
}

var map_PolicyBinding = map[string]string{
	"":             "PolicyBinding is a object that holds all the RoleBindings for a particular namespace.  There is one PolicyBinding document per referenced Policy namespace",
	"metadata":     "Standard object's metadata.",
	"lastModified": "LastModified is the last time that any part of the PolicyBinding was created, updated, or deleted",
	"policyRef":    "PolicyRef is a reference to the Policy that contains all the Roles that this PolicyBinding's RoleBindings may reference",
	"roleBindings": "RoleBindings holds all the RoleBindings held by this PolicyBinding, mapped by RoleBinding.Name",
}

func (PolicyBinding) SwaggerDoc() map[string]string {
	return map_PolicyBinding
}

var map_PolicyBindingList = map[string]string{
	"":         "PolicyBindingList is a collection of PolicyBindings",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of PolicyBindings",
}

func (PolicyBindingList) SwaggerDoc() map[string]string {
	return map_PolicyBindingList
}

var map_PolicyList = map[string]string{
	"":         "PolicyList is a collection of Policies",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of Policies",
}

func (PolicyList) SwaggerDoc() map[string]string {
	return map_PolicyList
}

var map_PolicyRule = map[string]string{
	"":                      "PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.",
	"verbs":                 "Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.",
	"attributeRestrictions": "AttributeRestrictions will vary depending on what the Authorizer/AuthorizationAttributeBuilder pair supports. If the Authorizer does not recognize how to handle the AttributeRestrictions, the Authorizer should report an error.",
	"apiGroups":             "APIGroups is the name of the APIGroup that contains the resources.  If this field is empty, then both kubernetes and origin API groups are assumed. That means that if an action is requested against one of the enumerated resources in either the kubernetes or the origin API group, the request will be allowed",
	"resources":             "Resources is a list of resources this rule applies to.  ResourceAll represents all resources.",
	"resourceNames":         "ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.",
	"nonResourceURLs":       "NonResourceURLsSlice is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path This name is intentionally different than the internal type so that the DefaultConvert works nicely and because the ordering may be different.",
}

func (PolicyRule) SwaggerDoc() map[string]string {
	return map_PolicyRule
}

var map_ResourceAccessReview = map[string]string{
	"": "ResourceAccessReview is a means to request a list of which users and groups are authorized to perform the action specified by spec",
}

func (ResourceAccessReview) SwaggerDoc() map[string]string {
	return map_ResourceAccessReview
}

var map_ResourceAccessReviewResponse = map[string]string{
	"":               "ResourceAccessReviewResponse describes who can perform the action",
	"namespace":      "Namespace is the namespace used for the access review",
	"users":          "UsersSlice is the list of users who can perform the action",
	"groups":         "GroupsSlice is the list of groups who can perform the action",
	"evalutionError": "EvaluationError is an indication that some error occurred during resolution, but partial results can still be returned. It is entirely possible to get an error and be able to continue determine authorization status in spite of it.  This is most common when a bound role is missing, but enough roles are still present and bound to reason about the request.",
}

func (ResourceAccessReviewResponse) SwaggerDoc() map[string]string {
	return map_ResourceAccessReviewResponse
}

var map_Role = map[string]string{
	"":         "Role is a logical grouping of PolicyRules that can be referenced as a unit by RoleBindings.",
	"metadata": "Standard object's metadata.",
	"rules":    "Rules holds all the PolicyRules for this Role",
}

func (Role) SwaggerDoc() map[string]string {
	return map_Role
}

var map_RoleBinding = map[string]string{
	"":           "RoleBinding references a Role, but not contain it.  It can reference any Role in the same namespace or in the global namespace. It adds who information via (Users and Groups) OR Subjects and namespace information by which namespace it exists in. RoleBindings in a given namespace only have effect in that namespace (excepting the master namespace which has power in all namespaces).",
	"metadata":   "Standard object's metadata.",
	"userNames":  "UserNames holds all the usernames directly bound to the role. This field should only be specified when supporting legacy clients and servers. See Subjects for further details.",
	"groupNames": "GroupNames holds all the groups directly bound to the role. This field should only be specified when supporting legacy clients and servers. See Subjects for further details.",
	"subjects":   "Subjects hold object references to authorize with this rule. This field is ignored if UserNames or GroupNames are specified to support legacy clients and servers. Thus newer clients that do not need to support backwards compatibility should send only fully qualified Subjects and should omit the UserNames and GroupNames fields. Clients that need to support backwards compatibility can use this field to build the UserNames and GroupNames.",
	"roleRef":    "RoleRef can only reference the current namespace and the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error. Since Policy is a singleton, this is sufficient knowledge to locate a role.",
}

func (RoleBinding) SwaggerDoc() map[string]string {
	return map_RoleBinding
}

var map_RoleBindingList = map[string]string{
	"":         "RoleBindingList is a collection of RoleBindings",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of RoleBindings",
}

func (RoleBindingList) SwaggerDoc() map[string]string {
	return map_RoleBindingList
}

var map_RoleBindingRestriction = map[string]string{
	"":         "RoleBindingRestriction is an object that can be matched against a subject (user, group, or service account) to determine whether rolebindings on that subject are allowed in the namespace to which the RoleBindingRestriction belongs.  If any one of those RoleBindingRestriction objects matches a subject, rolebindings on that subject in the namespace are allowed.",
	"metadata": "Standard object's metadata.",
	"spec":     "Spec defines the matcher.",
}

func (RoleBindingRestriction) SwaggerDoc() map[string]string {
	return map_RoleBindingRestriction
}

var map_RoleBindingRestrictionList = map[string]string{
	"":         "RoleBindingRestrictionList is a collection of RoleBindingRestriction objects.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of RoleBindingRestriction objects.",
}

func (RoleBindingRestrictionList) SwaggerDoc() map[string]string {
	return map_RoleBindingRestrictionList
}

var map_RoleBindingRestrictionSpec = map[string]string{
	"":                          "RoleBindingRestrictionSpec defines a rolebinding restriction.  Exactly one field must be non-nil.",
	"userrestriction":           "UserRestriction matches against user subjects.",
	"grouprestriction":          "GroupRestriction matches against group subjects.",
	"serviceaccountrestriction": "ServiceAccountRestriction matches against service-account subjects.",
}

func (RoleBindingRestrictionSpec) SwaggerDoc() map[string]string {
	return map_RoleBindingRestrictionSpec
}

var map_RoleList = map[string]string{
	"":         "RoleList is a collection of Roles",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of Roles",
}

func (RoleList) SwaggerDoc() map[string]string {
	return map_RoleList
}

var map_SelfSubjectRulesReview = map[string]string{
	"":       "SelfSubjectRulesReview is a resource you can create to determine which actions you can perform in a namespace",
	"spec":   "Spec adds information about how to conduct the check",
	"status": "Status is completed by the server to tell which permissions you have",
}

func (SelfSubjectRulesReview) SwaggerDoc() map[string]string {
	return map_SelfSubjectRulesReview
}

var map_SelfSubjectRulesReviewSpec = map[string]string{
	"":       "SelfSubjectRulesReviewSpec adds information about how to conduct the check",
	"scopes": "Scopes to use for the evaluation.  Empty means \"use the unscoped (full) permissions of the user/groups\". Nil means \"use the scopes on this request\".",
}

func (SelfSubjectRulesReviewSpec) SwaggerDoc() map[string]string {
	return map_SelfSubjectRulesReviewSpec
}

var map_ServiceAccountReference = map[string]string{
	"":          "ServiceAccountReference specifies a service account and namespace by their names.",
	"name":      "Name is the name of the service account.",
	"namespace": "Namespace is the namespace of the service account.  Service accounts from inside the whitelisted namespaces are allowed to be bound to roles.  If Namespace is empty, then the namespace of the RoleBindingRestriction in which the ServiceAccountReference is embedded is used.",
}

func (ServiceAccountReference) SwaggerDoc() map[string]string {
	return map_ServiceAccountReference
}

var map_ServiceAccountRestriction = map[string]string{
	"":                "ServiceAccountRestriction matches a service account by a string match on either the service-account name or the name of the service account's namespace.",
	"serviceaccounts": "ServiceAccounts specifies a list of literal service-account names.",
	"namespaces":      "Namespaces specifies a list of literal namespace names.",
}

func (ServiceAccountRestriction) SwaggerDoc() map[string]string {
	return map_ServiceAccountRestriction
}

var map_SubjectAccessReview = map[string]string{
	"":       "SubjectAccessReview is an object for requesting information about whether a user or group can perform an action",
	"user":   "User is optional. If both User and Groups are empty, the current authenticated user is used.",
	"groups": "GroupsSlice is optional. Groups is the list of groups to which the User belongs.",
	"scopes": "Scopes to use for the evaluation.  Empty means \"use the unscoped (full) permissions of the user/groups\". Nil for a self-SAR, means \"use the scopes on this request\". Nil for a regular SAR, means the same as empty.",
}

func (SubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SubjectAccessReview
}

var map_SubjectAccessReviewResponse = map[string]string{
	"":                "SubjectAccessReviewResponse describes whether or not a user or group can perform an action",
	"namespace":       "Namespace is the namespace used for the access review",
	"allowed":         "Allowed is required.  True if the action would be allowed, false otherwise.",
	"reason":          "Reason is optional.  It indicates why a request was allowed or denied.",
	"evaluationError": "EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it.  This is most common when a bound role is missing, but enough roles are still present and bound to reason about the request.",
}

func (SubjectAccessReviewResponse) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewResponse
}

var map_SubjectRulesReview = map[string]string{
	"":       "SubjectRulesReview is a resource you can create to determine which actions another user can perform in a namespace",
	"spec":   "Spec adds information about how to conduct the check",
	"status": "Status is completed by the server to tell which permissions you have",
}

func (SubjectRulesReview) SwaggerDoc() map[string]string {
	return map_SubjectRulesReview
}

var map_SubjectRulesReviewSpec = map[string]string{
	"":       "SubjectRulesReviewSpec adds information about how to conduct the check",
	"user":   "User is optional.  At least one of User and Groups must be specified.",
	"groups": "Groups is optional.  Groups is the list of groups to which the User belongs.  At least one of User and Groups must be specified.",
	"scopes": "Scopes to use for the evaluation.  Empty means \"use the unscoped (full) permissions of the user/groups\".",
}

func (SubjectRulesReviewSpec) SwaggerDoc() map[string]string {
	return map_SubjectRulesReviewSpec
}

var map_SubjectRulesReviewStatus = map[string]string{
	"":                "SubjectRulesReviewStatus is contains the result of a rules check",
	"rules":           "Rules is the list of rules (no particular sort) that are allowed for the subject",
	"evaluationError": "EvaluationError can appear in combination with Rules.  It means some error happened during evaluation that may have prevented additional rules from being populated.",
}

func (SubjectRulesReviewStatus) SwaggerDoc() map[string]string {
	return map_SubjectRulesReviewStatus
}

var map_UserRestriction = map[string]string{
	"":       "UserRestriction matches a user either by a string match on the user name, a string match on the name of a group to which the user belongs, or a label selector applied to the user labels.",
	"users":  "Users specifies a list of literal user names.",
	"groups": "Groups specifies a list of literal group names.",
	"labels": "Selectors specifies a list of label selectors over user labels.",
}

func (UserRestriction) SwaggerDoc() map[string]string {
	return map_UserRestriction
}
