// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/duc-cnzj/mars/api/v4/types"
	"github.com/duc-cnzj/mars/v4/internal/ent/accesstoken"
	"github.com/duc-cnzj/mars/v4/internal/ent/cachelock"
	"github.com/duc-cnzj/mars/v4/internal/ent/changelog"
	"github.com/duc-cnzj/mars/v4/internal/ent/dbcache"
	"github.com/duc-cnzj/mars/v4/internal/ent/event"
	"github.com/duc-cnzj/mars/v4/internal/ent/file"
	"github.com/duc-cnzj/mars/v4/internal/ent/gitproject"
	"github.com/duc-cnzj/mars/v4/internal/ent/namespace"
	"github.com/duc-cnzj/mars/v4/internal/ent/project"
	"github.com/duc-cnzj/mars/v4/internal/ent/schema"
	"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accesstokenMixin := schema.AccessToken{}.Mixin()
	accesstokenMixinHooks2 := accesstokenMixin[2].Hooks()
	accesstoken.Hooks[0] = accesstokenMixinHooks2[0]
	accesstokenMixinInters2 := accesstokenMixin[2].Interceptors()
	accesstoken.Interceptors[0] = accesstokenMixinInters2[0]
	accesstokenMixinFields0 := accesstokenMixin[0].Fields()
	_ = accesstokenMixinFields0
	accesstokenMixinFields1 := accesstokenMixin[1].Fields()
	_ = accesstokenMixinFields1
	accesstokenFields := schema.AccessToken{}.Fields()
	_ = accesstokenFields
	// accesstokenDescCreatedAt is the schema descriptor for created_at field.
	accesstokenDescCreatedAt := accesstokenMixinFields0[0].Descriptor()
	// accesstoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	accesstoken.DefaultCreatedAt = accesstokenDescCreatedAt.Default.(func() time.Time)
	// accesstokenDescUpdatedAt is the schema descriptor for updated_at field.
	accesstokenDescUpdatedAt := accesstokenMixinFields1[0].Descriptor()
	// accesstoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	accesstoken.DefaultUpdatedAt = accesstokenDescUpdatedAt.Default.(func() time.Time)
	// accesstoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	accesstoken.UpdateDefaultUpdatedAt = accesstokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accesstokenDescToken is the schema descriptor for token field.
	accesstokenDescToken := accesstokenFields[0].Descriptor()
	// accesstoken.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	accesstoken.TokenValidator = func() func(string) error {
		validators := accesstokenDescToken.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(token string) error {
			for _, fn := range fns {
				if err := fn(token); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// accesstokenDescUsage is the schema descriptor for usage field.
	accesstokenDescUsage := accesstokenFields[1].Descriptor()
	// accesstoken.UsageValidator is a validator for the "usage" field. It is called by the builders before save.
	accesstoken.UsageValidator = accesstokenDescUsage.Validators[0].(func(string) error)
	// accesstokenDescEmail is the schema descriptor for email field.
	accesstokenDescEmail := accesstokenFields[2].Descriptor()
	// accesstoken.DefaultEmail holds the default value on creation for the email field.
	accesstoken.DefaultEmail = accesstokenDescEmail.Default.(string)
	// accesstoken.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	accesstoken.EmailValidator = accesstokenDescEmail.Validators[0].(func(string) error)
	cachelockFields := schema.CacheLock{}.Fields()
	_ = cachelockFields
	// cachelockDescKey is the schema descriptor for key field.
	cachelockDescKey := cachelockFields[0].Descriptor()
	// cachelock.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	cachelock.KeyValidator = cachelockDescKey.Validators[0].(func(string) error)
	// cachelockDescOwner is the schema descriptor for owner field.
	cachelockDescOwner := cachelockFields[1].Descriptor()
	// cachelock.OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	cachelock.OwnerValidator = cachelockDescOwner.Validators[0].(func(string) error)
	changelogMixin := schema.Changelog{}.Mixin()
	changelogMixinHooks2 := changelogMixin[2].Hooks()
	changelog.Hooks[0] = changelogMixinHooks2[0]
	changelogMixinInters2 := changelogMixin[2].Interceptors()
	changelog.Interceptors[0] = changelogMixinInters2[0]
	changelogMixinFields0 := changelogMixin[0].Fields()
	_ = changelogMixinFields0
	changelogMixinFields1 := changelogMixin[1].Fields()
	_ = changelogMixinFields1
	changelogFields := schema.Changelog{}.Fields()
	_ = changelogFields
	// changelogDescCreatedAt is the schema descriptor for created_at field.
	changelogDescCreatedAt := changelogMixinFields0[0].Descriptor()
	// changelog.DefaultCreatedAt holds the default value on creation for the created_at field.
	changelog.DefaultCreatedAt = changelogDescCreatedAt.Default.(func() time.Time)
	// changelogDescUpdatedAt is the schema descriptor for updated_at field.
	changelogDescUpdatedAt := changelogMixinFields1[0].Descriptor()
	// changelog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	changelog.DefaultUpdatedAt = changelogDescUpdatedAt.Default.(func() time.Time)
	// changelog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	changelog.UpdateDefaultUpdatedAt = changelogDescUpdatedAt.UpdateDefault.(func() time.Time)
	// changelogDescVersion is the schema descriptor for version field.
	changelogDescVersion := changelogFields[0].Descriptor()
	// changelog.DefaultVersion holds the default value on creation for the version field.
	changelog.DefaultVersion = changelogDescVersion.Default.(int)
	// changelogDescUsername is the schema descriptor for username field.
	changelogDescUsername := changelogFields[1].Descriptor()
	// changelog.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	changelog.UsernameValidator = changelogDescUsername.Validators[0].(func(string) error)
	// changelogDescConfigType is the schema descriptor for config_type field.
	changelogDescConfigType := changelogFields[4].Descriptor()
	// changelog.ConfigTypeValidator is a validator for the "config_type" field. It is called by the builders before save.
	changelog.ConfigTypeValidator = changelogDescConfigType.Validators[0].(func(string) error)
	// changelogDescGitBranch is the schema descriptor for git_branch field.
	changelogDescGitBranch := changelogFields[5].Descriptor()
	// changelog.GitBranchValidator is a validator for the "git_branch" field. It is called by the builders before save.
	changelog.GitBranchValidator = func() func(string) error {
		validators := changelogDescGitBranch.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(git_branch string) error {
			for _, fn := range fns {
				if err := fn(git_branch); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// changelogDescGitCommit is the schema descriptor for git_commit field.
	changelogDescGitCommit := changelogFields[6].Descriptor()
	// changelog.GitCommitValidator is a validator for the "git_commit" field. It is called by the builders before save.
	changelog.GitCommitValidator = func() func(string) error {
		validators := changelogDescGitCommit.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(git_commit string) error {
			for _, fn := range fns {
				if err := fn(git_commit); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// changelogDescGitCommitWebURL is the schema descriptor for git_commit_web_url field.
	changelogDescGitCommitWebURL := changelogFields[11].Descriptor()
	// changelog.GitCommitWebURLValidator is a validator for the "git_commit_web_url" field. It is called by the builders before save.
	changelog.GitCommitWebURLValidator = changelogDescGitCommitWebURL.Validators[0].(func(string) error)
	// changelogDescGitCommitTitle is the schema descriptor for git_commit_title field.
	changelogDescGitCommitTitle := changelogFields[12].Descriptor()
	// changelog.GitCommitTitleValidator is a validator for the "git_commit_title" field. It is called by the builders before save.
	changelog.GitCommitTitleValidator = changelogDescGitCommitTitle.Validators[0].(func(string) error)
	// changelogDescGitCommitAuthor is the schema descriptor for git_commit_author field.
	changelogDescGitCommitAuthor := changelogFields[13].Descriptor()
	// changelog.GitCommitAuthorValidator is a validator for the "git_commit_author" field. It is called by the builders before save.
	changelog.GitCommitAuthorValidator = changelogDescGitCommitAuthor.Validators[0].(func(string) error)
	// changelogDescConfigChanged is the schema descriptor for config_changed field.
	changelogDescConfigChanged := changelogFields[15].Descriptor()
	// changelog.DefaultConfigChanged holds the default value on creation for the config_changed field.
	changelog.DefaultConfigChanged = changelogDescConfigChanged.Default.(bool)
	dbcacheFields := schema.DBCache{}.Fields()
	_ = dbcacheFields
	// dbcacheDescKey is the schema descriptor for key field.
	dbcacheDescKey := dbcacheFields[0].Descriptor()
	// dbcache.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	dbcache.KeyValidator = dbcacheDescKey.Validators[0].(func(string) error)
	eventMixin := schema.Event{}.Mixin()
	eventMixinHooks2 := eventMixin[2].Hooks()
	event.Hooks[0] = eventMixinHooks2[0]
	eventMixinInters2 := eventMixin[2].Interceptors()
	event.Interceptors[0] = eventMixinInters2[0]
	eventMixinFields0 := eventMixin[0].Fields()
	_ = eventMixinFields0
	eventMixinFields1 := eventMixin[1].Fields()
	_ = eventMixinFields1
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventMixinFields0[0].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventMixinFields1[0].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	// eventDescAction is the schema descriptor for action field.
	eventDescAction := eventFields[0].Descriptor()
	// event.DefaultAction holds the default value on creation for the action field.
	event.DefaultAction = types.EventActionType(eventDescAction.Default.(int32))
	// eventDescUsername is the schema descriptor for username field.
	eventDescUsername := eventFields[1].Descriptor()
	// event.DefaultUsername holds the default value on creation for the username field.
	event.DefaultUsername = eventDescUsername.Default.(string)
	// event.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	event.UsernameValidator = eventDescUsername.Validators[0].(func(string) error)
	// eventDescMessage is the schema descriptor for message field.
	eventDescMessage := eventFields[2].Descriptor()
	// event.DefaultMessage holds the default value on creation for the message field.
	event.DefaultMessage = eventDescMessage.Default.(string)
	// event.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	event.MessageValidator = eventDescMessage.Validators[0].(func(string) error)
	// eventDescDuration is the schema descriptor for duration field.
	eventDescDuration := eventFields[5].Descriptor()
	// event.DefaultDuration holds the default value on creation for the duration field.
	event.DefaultDuration = eventDescDuration.Default.(string)
	fileMixin := schema.File{}.Mixin()
	fileMixinHooks2 := fileMixin[2].Hooks()
	file.Hooks[0] = fileMixinHooks2[0]
	fileMixinInters2 := fileMixin[2].Interceptors()
	file.Interceptors[0] = fileMixinInters2[0]
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileMixinFields1 := fileMixin[1].Fields()
	_ = fileMixinFields1
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileMixinFields0[0].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileMixinFields1[0].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescUploadType is the schema descriptor for upload_type field.
	fileDescUploadType := fileFields[0].Descriptor()
	// file.DefaultUploadType holds the default value on creation for the upload_type field.
	file.DefaultUploadType = schematype.UploadType(fileDescUploadType.Default.(string))
	// file.UploadTypeValidator is a validator for the "upload_type" field. It is called by the builders before save.
	file.UploadTypeValidator = fileDescUploadType.Validators[0].(func(string) error)
	// fileDescPath is the schema descriptor for path field.
	fileDescPath := fileFields[1].Descriptor()
	// file.PathValidator is a validator for the "path" field. It is called by the builders before save.
	file.PathValidator = fileDescPath.Validators[0].(func(string) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[2].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(uint64)
	// fileDescUsername is the schema descriptor for username field.
	fileDescUsername := fileFields[3].Descriptor()
	// file.DefaultUsername holds the default value on creation for the username field.
	file.DefaultUsername = fileDescUsername.Default.(string)
	// file.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	file.UsernameValidator = fileDescUsername.Validators[0].(func(string) error)
	// fileDescNamespace is the schema descriptor for namespace field.
	fileDescNamespace := fileFields[4].Descriptor()
	// file.DefaultNamespace holds the default value on creation for the namespace field.
	file.DefaultNamespace = fileDescNamespace.Default.(string)
	// file.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	file.NamespaceValidator = fileDescNamespace.Validators[0].(func(string) error)
	// fileDescPod is the schema descriptor for pod field.
	fileDescPod := fileFields[5].Descriptor()
	// file.DefaultPod holds the default value on creation for the pod field.
	file.DefaultPod = fileDescPod.Default.(string)
	// file.PodValidator is a validator for the "pod" field. It is called by the builders before save.
	file.PodValidator = fileDescPod.Validators[0].(func(string) error)
	// fileDescContainer is the schema descriptor for container field.
	fileDescContainer := fileFields[6].Descriptor()
	// file.DefaultContainer holds the default value on creation for the container field.
	file.DefaultContainer = fileDescContainer.Default.(string)
	// file.ContainerValidator is a validator for the "container" field. It is called by the builders before save.
	file.ContainerValidator = fileDescContainer.Validators[0].(func(string) error)
	// fileDescContainerPath is the schema descriptor for container_path field.
	fileDescContainerPath := fileFields[7].Descriptor()
	// file.DefaultContainerPath holds the default value on creation for the container_path field.
	file.DefaultContainerPath = fileDescContainerPath.Default.(string)
	// file.ContainerPathValidator is a validator for the "container_path" field. It is called by the builders before save.
	file.ContainerPathValidator = fileDescContainerPath.Validators[0].(func(string) error)
	gitprojectMixin := schema.GitProject{}.Mixin()
	gitprojectMixinHooks2 := gitprojectMixin[2].Hooks()
	gitproject.Hooks[0] = gitprojectMixinHooks2[0]
	gitprojectMixinInters2 := gitprojectMixin[2].Interceptors()
	gitproject.Interceptors[0] = gitprojectMixinInters2[0]
	gitprojectMixinFields0 := gitprojectMixin[0].Fields()
	_ = gitprojectMixinFields0
	gitprojectMixinFields1 := gitprojectMixin[1].Fields()
	_ = gitprojectMixinFields1
	gitprojectFields := schema.GitProject{}.Fields()
	_ = gitprojectFields
	// gitprojectDescCreatedAt is the schema descriptor for created_at field.
	gitprojectDescCreatedAt := gitprojectMixinFields0[0].Descriptor()
	// gitproject.DefaultCreatedAt holds the default value on creation for the created_at field.
	gitproject.DefaultCreatedAt = gitprojectDescCreatedAt.Default.(func() time.Time)
	// gitprojectDescUpdatedAt is the schema descriptor for updated_at field.
	gitprojectDescUpdatedAt := gitprojectMixinFields1[0].Descriptor()
	// gitproject.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	gitproject.DefaultUpdatedAt = gitprojectDescUpdatedAt.Default.(func() time.Time)
	// gitproject.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	gitproject.UpdateDefaultUpdatedAt = gitprojectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// gitprojectDescName is the schema descriptor for name field.
	gitprojectDescName := gitprojectFields[0].Descriptor()
	// gitproject.DefaultName holds the default value on creation for the name field.
	gitproject.DefaultName = gitprojectDescName.Default.(string)
	// gitproject.NameValidator is a validator for the "name" field. It is called by the builders before save.
	gitproject.NameValidator = gitprojectDescName.Validators[0].(func(string) error)
	// gitprojectDescDefaultBranch is the schema descriptor for default_branch field.
	gitprojectDescDefaultBranch := gitprojectFields[1].Descriptor()
	// gitproject.DefaultDefaultBranch holds the default value on creation for the default_branch field.
	gitproject.DefaultDefaultBranch = gitprojectDescDefaultBranch.Default.(string)
	// gitproject.DefaultBranchValidator is a validator for the "default_branch" field. It is called by the builders before save.
	gitproject.DefaultBranchValidator = gitprojectDescDefaultBranch.Validators[0].(func(string) error)
	// gitprojectDescEnabled is the schema descriptor for enabled field.
	gitprojectDescEnabled := gitprojectFields[3].Descriptor()
	// gitproject.DefaultEnabled holds the default value on creation for the enabled field.
	gitproject.DefaultEnabled = gitprojectDescEnabled.Default.(bool)
	// gitprojectDescGlobalEnabled is the schema descriptor for global_enabled field.
	gitprojectDescGlobalEnabled := gitprojectFields[4].Descriptor()
	// gitproject.DefaultGlobalEnabled holds the default value on creation for the global_enabled field.
	gitproject.DefaultGlobalEnabled = gitprojectDescGlobalEnabled.Default.(bool)
	namespaceMixin := schema.Namespace{}.Mixin()
	namespaceMixinHooks2 := namespaceMixin[2].Hooks()
	namespace.Hooks[0] = namespaceMixinHooks2[0]
	namespaceMixinInters2 := namespaceMixin[2].Interceptors()
	namespace.Interceptors[0] = namespaceMixinInters2[0]
	namespaceMixinFields0 := namespaceMixin[0].Fields()
	_ = namespaceMixinFields0
	namespaceMixinFields1 := namespaceMixin[1].Fields()
	_ = namespaceMixinFields1
	namespaceFields := schema.Namespace{}.Fields()
	_ = namespaceFields
	// namespaceDescCreatedAt is the schema descriptor for created_at field.
	namespaceDescCreatedAt := namespaceMixinFields0[0].Descriptor()
	// namespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	namespace.DefaultCreatedAt = namespaceDescCreatedAt.Default.(func() time.Time)
	// namespaceDescUpdatedAt is the schema descriptor for updated_at field.
	namespaceDescUpdatedAt := namespaceMixinFields1[0].Descriptor()
	// namespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	namespace.DefaultUpdatedAt = namespaceDescUpdatedAt.Default.(func() time.Time)
	// namespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	namespace.UpdateDefaultUpdatedAt = namespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// namespaceDescName is the schema descriptor for name field.
	namespaceDescName := namespaceFields[0].Descriptor()
	// namespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	namespace.NameValidator = namespaceDescName.Validators[0].(func(string) error)
	// namespaceDescImagePullSecrets is the schema descriptor for image_pull_secrets field.
	namespaceDescImagePullSecrets := namespaceFields[1].Descriptor()
	// namespace.DefaultImagePullSecrets holds the default value on creation for the image_pull_secrets field.
	namespace.DefaultImagePullSecrets = namespaceDescImagePullSecrets.Default.([]string)
	projectMixin := schema.Project{}.Mixin()
	projectMixinHooks2 := projectMixin[2].Hooks()
	project.Hooks[0] = projectMixinHooks2[0]
	projectMixinInters2 := projectMixin[2].Interceptors()
	project.Interceptors[0] = projectMixinInters2[0]
	projectMixinFields0 := projectMixin[0].Fields()
	_ = projectMixinFields0
	projectMixinFields1 := projectMixin[1].Fields()
	_ = projectMixinFields1
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectMixinFields0[0].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectMixinFields1[0].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[0].Descriptor()
	// project.DefaultName holds the default value on creation for the name field.
	project.DefaultName = projectDescName.Default.(string)
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescGitBranch is the schema descriptor for git_branch field.
	projectDescGitBranch := projectFields[2].Descriptor()
	// project.GitBranchValidator is a validator for the "git_branch" field. It is called by the builders before save.
	project.GitBranchValidator = projectDescGitBranch.Validators[0].(func(string) error)
	// projectDescGitCommit is the schema descriptor for git_commit field.
	projectDescGitCommit := projectFields[3].Descriptor()
	// project.GitCommitValidator is a validator for the "git_commit" field. It is called by the builders before save.
	project.GitCommitValidator = projectDescGitCommit.Validators[0].(func(string) error)
	// projectDescAtomic is the schema descriptor for atomic field.
	projectDescAtomic := projectFields[8].Descriptor()
	// project.DefaultAtomic holds the default value on creation for the atomic field.
	project.DefaultAtomic = projectDescAtomic.Default.(bool)
	// projectDescDeployStatus is the schema descriptor for deploy_status field.
	projectDescDeployStatus := projectFields[9].Descriptor()
	// project.DefaultDeployStatus holds the default value on creation for the deploy_status field.
	project.DefaultDeployStatus = types.Deploy(projectDescDeployStatus.Default.(int32))
	// projectDescVersion is the schema descriptor for version field.
	projectDescVersion := projectFields[13].Descriptor()
	// project.DefaultVersion holds the default value on creation for the version field.
	project.DefaultVersion = projectDescVersion.Default.(int)
	// projectDescConfigType is the schema descriptor for config_type field.
	projectDescConfigType := projectFields[14].Descriptor()
	// project.ConfigTypeValidator is a validator for the "config_type" field. It is called by the builders before save.
	project.ConfigTypeValidator = projectDescConfigType.Validators[0].(func(string) error)
	// projectDescGitCommitWebURL is the schema descriptor for git_commit_web_url field.
	projectDescGitCommitWebURL := projectFields[16].Descriptor()
	// project.DefaultGitCommitWebURL holds the default value on creation for the git_commit_web_url field.
	project.DefaultGitCommitWebURL = projectDescGitCommitWebURL.Default.(string)
	// project.GitCommitWebURLValidator is a validator for the "git_commit_web_url" field. It is called by the builders before save.
	project.GitCommitWebURLValidator = projectDescGitCommitWebURL.Validators[0].(func(string) error)
	// projectDescGitCommitTitle is the schema descriptor for git_commit_title field.
	projectDescGitCommitTitle := projectFields[17].Descriptor()
	// project.DefaultGitCommitTitle holds the default value on creation for the git_commit_title field.
	project.DefaultGitCommitTitle = projectDescGitCommitTitle.Default.(string)
	// project.GitCommitTitleValidator is a validator for the "git_commit_title" field. It is called by the builders before save.
	project.GitCommitTitleValidator = projectDescGitCommitTitle.Validators[0].(func(string) error)
	// projectDescGitCommitAuthor is the schema descriptor for git_commit_author field.
	projectDescGitCommitAuthor := projectFields[18].Descriptor()
	// project.DefaultGitCommitAuthor holds the default value on creation for the git_commit_author field.
	project.DefaultGitCommitAuthor = projectDescGitCommitAuthor.Default.(string)
	// project.GitCommitAuthorValidator is a validator for the "git_commit_author" field. It is called by the builders before save.
	project.GitCommitAuthorValidator = projectDescGitCommitAuthor.Validators[0].(func(string) error)
}

const (
	Version = "v0.14.0"                                         // Version of ent codegen.
	Sum     = "h1:EO3Z9aZ5bXJatJeGqu/EVdnNr6K4mRq3rWe5owt0MC4=" // Sum of ent codegen.
)
