package v2

var Commands commands

type commands struct {
	App                                AppCommand                                `command:"app"`
	Help                               HelpCommand                               `command:"help"`
	Version                            VersionCommand                            `command:"version"`
	Login                              LoginCommand                              `command:"login"`
	Logout                             LogoutCommand                             `command:"logout"`
	Passwd                             PasswdCommand                             `command:"passwd"`
	Target                             TargetCommand                             `command:"target"`
	Api                                ApiCommand                                `command:"api"`
	Auth                               AuthCommand                               `command:"auth"`
	Apps                               AppsCommand                               `command:"apps"`
	Push                               PushCommand                               `command:"push"`
	Scale                              ScaleCommand                              `command:"scale"`
	Delete                             DeleteCommand                             `command:"delete"`
	Rename                             RenameCommand                             `command:"rename"`
	Start                              StartCommand                              `command:"start"`
	Stop                               StopCommand                               `command:"stop"`
	Restart                            RestartCommand                            `command:"restart"`
	Restage                            RestageCommand                            `command:"restage"`
	RestartAppInstance                 RestartAppInstanceCommand                 `command:"restart-app-instance"`
	Events                             EventsCommand                             `command:"events"`
	Files                              FilesCommand                              `command:"files"`
	Logs                               LogsCommand                               `command:"logs"`
	Env                                EnvCommand                                `command:"env"`
	SetEnv                             SetEnvCommand                             `command:"set-env"`
	UnsetEnv                           UnsetEnvCommand                           `command:"unset-env"`
	Stacks                             StacksCommand                             `command:"stacks"`
	Stack                              StackCommand                              `command:"stack"`
	CopySource                         CopySourceCommand                         `command:"copy-source"`
	CreateAppManifest                  CreateAppManifestCommand                  `command:"create-app-manifest"`
	GetHealthCheck                     GetHealthCheckCommand                     `command:"get-health-check"`
	SetHealthCheck                     SetHealthCheckCommand                     `command:"set-health-check"`
	EnableSSH                          EnableSSHCommand                          `command:"enable-ssh"`
	DisableSSH                         DisableSSHCommand                         `command:"disable-ssh"`
	SSHEnabled                         SSHEnabledCommand                         `command:"ssh-enabled"`
	SSH                                SSHCommand                                `command:"ssh"`
	Marketplace                        MarketplaceCommand                        `command:"marketplace"`
	Services                           ServicesCommand                           `command:"services"`
	Service                            ServiceCommand                            `command:"service"`
	CreateService                      CreateServiceCommand                      `command:"create-service"`
	UpdateService                      UpdateServiceCommand                      `command:"update-service"`
	DeleteService                      DeleteServiceCommand                      `command:"delete-service"`
	RenameService                      RenameServiceCommand                      `command:"rename-service"`
	CreateServiceKey                   CreateServiceKeyCommand                   `command:"create-service-key"`
	ServiceKeys                        ServiceKeysCommand                        `command:"service-keys"`
	ServiceKey                         ServiceKeyCommand                         `command:"service-key"`
	DeleteServiceKey                   DeleteServiceKeyCommand                   `command:"delete-service-key"`
	BindService                        BindServiceCommand                        `command:"bind-service"`
	UnbindService                      UnbindServiceCommand                      `command:"unbind-service"`
	BindRouteService                   BindRouteServiceCommand                   `command:"bind-route-service"`
	UnbindRouteService                 UnbindRouteServiceCommand                 `command:"unbind-route-service"`
	CreateUserProvidedService          CreateUserProvidedServiceCommand          `command:"create-user-provided-service"`
	UpdateUserProvidedService          UpdateUserProvidedServiceCommand          `command:"update-user-provided-service"`
	Orgs                               OrgsCommand                               `command:"orgs"`
	Org                                OrgCommand                                `command:"org"`
	CreateOrg                          CreateOrgCommand                          `command:"create-org"`
	DeleteOrg                          DeleteOrgCommand                          `command:"delete-org"`
	RenameOrg                          RenameOrgCommand                          `command:"rename-org"`
	Spaces                             SpacesCommand                             `command:"spaces"`
	Space                              SpaceCommand                              `command:"space"`
	CreateSpace                        CreateSpaceCommand                        `command:"create-space"`
	DeleteSpace                        DeleteSpaceCommand                        `command:"delete-space"`
	RenameSpace                        RenameSpaceCommand                        `command:"rename-space"`
	AllowSpaceSSH                      AllowSpaceSSHCommand                      `command:"allow-space-ssh"`
	DisallowSpaceSSH                   DisallowSpaceSSHCommand                   `command:"disallow-space-ssh"`
	SpaceSSHAllowed                    SpaceSSHAllowedCommand                    `command:"space-ssh-allowed"`
	Domains                            DomainsCommand                            `command:"domains"`
	CreateDomain                       CreateDomainCommand                       `command:"create-domain"`
	DeleteDomain                       DeleteDomainCommand                       `command:"delete-domain"`
	CreateSharedDomain                 CreateSharedDomainCommand                 `command:"create-shared-domain"`
	DeleteSharedDomain                 DeleteSharedDomainCommand                 `command:"delete-shared-domain"`
	RouterGroups                       RouterGroupsCommand                       `command:"router-groups"`
	Routes                             RoutesCommand                             `command:"routes"`
	CreateRoute                        CreateRouteCommand                        `command:"create-route"`
	CheckRoute                         CheckRouteCommand                         `command:"check-route"`
	MapRoute                           MapRouteCommand                           `command:"map-route"`
	UnmapRoute                         UnmapRouteCommand                         `command:"unmap-route"`
	DeleteRoute                        DeleteRouteCommand                        `command:"delete-route"`
	DeleteOrphanedRoutes               DeleteOrphanedRoutesCommand               `command:"delete-orphaned-routes"`
	Buildpacks                         BuildpacksCommand                         `command:"buildpacks"`
	CreateBuildpack                    CreateBuildpackCommand                    `command:"create-buildpack"`
	UpdateBuildpack                    UpdateBuildpackCommand                    `command:"update-buildpack"`
	RenameBuildpack                    RenameBuildpackCommand                    `command:"rename-buildpack"`
	DeleteBuildpack                    DeleteBuildpackCommand                    `command:"delete-buildpack"`
	CreateUser                         CreateUserCommand                         `command:"create-user"`
	DeleteUser                         DeleteUserCommand                         `command:"delete-user"`
	OrgUsers                           OrgUsersCommand                           `command:"org-users"`
	SetOrgRole                         SetOrgRoleCommand                         `command:"set-org-role"`
	UnsetOrgRole                       UnsetOrgRoleCommand                       `command:"unset-org-role"`
	SpaceUsers                         SpaceUsersCommand                         `command:"space-users"`
	SetSpaceRole                       SetSpaceRoleCommand                       `command:"set-space-role"`
	UnsetSpaceRole                     UnsetSpaceRoleCommand                     `command:"unset-space-role"`
	Quotas                             QuotasCommand                             `command:"quotas"`
	Quota                              QuotaCommand                              `command:"quota"`
	SetQuota                           SetQuotaCommand                           `command:"set-quota"`
	CreateQuota                        CreateQuotaCommand                        `command:"create-quota"`
	DeleteQuota                        DeleteQuotaCommand                        `command:"delete-quota"`
	UpdateQuota                        UpdateQuotaCommand                        `command:"update-quota"`
	SharePrivateDomain                 SharePrivateDomainCommand                 `command:"share-private-domain"`
	UnsharePrivateDomain               UnsharePrivateDomainCommand               `command:"unshare-private-domain"`
	SpaceQuotas                        SpaceQuotasCommand                        `command:"space-quotas"`
	SpaceQuota                         SpaceQuotaCommand                         `command:"space-quota"`
	CreateSpaceQuota                   CreateSpaceQuotaCommand                   `command:"create-space-quota"`
	UpdateSpaceQuota                   UpdateSpaceQuotaCommand                   `command:"update-space-quota"`
	DeleteSpaceQuota                   DeleteSpaceQuotaCommand                   `command:"delete-space-quota"`
	SetSpaceQuota                      SetSpaceQuotaCommand                      `command:"set-space-quota"`
	UnsetSpaceQuota                    UnsetSpaceQuotaCommand                    `command:"unset-space-quota"`
	ServiceAuthTokens                  ServiceAuthTokensCommand                  `command:"service-auth-tokens"`
	CreateServiceAuthToken             CreateServiceAuthTokenCommand             `command:"create-service-auth-token"`
	UpdateServiceAuthToken             UpdateServiceAuthTokenCommand             `command:"update-service-auth-token"`
	DeleteServiceAuthToken             DeleteServiceAuthTokenCommand             `command:"delete-service-auth-token"`
	ServiceBrokers                     ServiceBrokersCommand                     `command:"service-brokers"`
	CreateServiceBroker                CreateServiceBrokerCommand                `command:"create-service-broker"`
	UpdateServiceBroker                UpdateServiceBrokerCommand                `command:"update-service-broker"`
	DeleteServiceBroker                DeleteServiceBrokerCommand                `command:"delete-service-broker"`
	RenameServiceBroker                RenameServiceBrokerCommand                `command:"rename-service-broker"`
	MigrateServiceInstances            MigrateServiceInstancesCommand            `command:"migrate-service-instances"`
	PurgeServiceOffering               PurgeServiceOfferingCommand               `command:"purge-service-offering"`
	PurgeServiceInstance               PurgeServiceInstanceCommand               `command:"purge-service-instance"`
	ServiceAccess                      ServiceAccessCommand                      `command:"service-access"`
	EnableServiceAccess                EnableServiceAccessCommand                `command:"enable-service-access"`
	DisableServiceAccess               DisableServiceAccessCommand               `command:"disable-service-access"`
	SecurityGroup                      SecurityGroupCommand                      `command:"security-group"`
	SecurityGroups                     SecurityGroupsCommand                     `command:"security-groups"`
	CreateSecurityGroup                CreateSecurityGroupCommand                `command:"create-security-group"`
	UpdateSecurityGroup                UpdateSecurityGroupCommand                `command:"update-security-group"`
	DeleteSecurityGroup                DeleteSecurityGroupCommand                `command:"delete-security-group"`
	BindSecurityGroup                  BindSecurityGroupCommand                  `command:"bind-security-group"`
	UnbindSecurityGroup                UnbindSecurityGroupCommand                `command:"unbind-security-group"`
	BindStagingSecurityGroup           BindStagingSecurityGroupCommand           `command:"bind-staging-security-group"`
	StagingSecurityGroups              StagingSecurityGroupsCommand              `command:"staging-security-groups"`
	UnbindStagingSecurityGroup         UnbindStagingSecurityGroupCommand         `command:"unbind-staging-security-group"`
	BindRunningSecurityGroup           BindRunningSecurityGroupCommand           `command:"bind-running-security-group"`
	RunningSecurityGroups              RunningSecurityGroupsCommand              `command:"running-security-groups"`
	UnbindRunningSecurityGroup         UnbindRunningSecurityGroupCommand         `command:"unbind-running-security-group"`
	RunningEnvironmentVariableGroup    RunningEnvironmentVariableGroupCommand    `command:"running-environment-variable-group"`
	StagingEnvironmentVariableGroup    StagingEnvironmentVariableGroupCommand    `command:"staging-environment-variable-group"`
	SetStagingEnvironmentVariableGroup SetStagingEnvironmentVariableGroupCommand `command:"set-staging-environment-variable-group"`
	SetRunningEnvironmentVariableGroup SetRunningEnvironmentVariableGroupCommand `command:"set-running-environment-variable-group"`
	FeatureFlags                       FeatureFlagsCommand                       `command:"feature-flags"`
	FeatureFlag                        FeatureFlagCommand                        `command:"feature-flag"`
	EnableFeatureFlag                  EnableFeatureFlagCommand                  `command:"enable-feature-flag"`
	DisableFeatureFlag                 DisableFeatureFlagCommand                 `command:"disable-feature-flag"`
	Curl                               CurlCommand                               `command:"curl"`
	Config                             ConfigCommand                             `command:"config"`
	OauthToken                         OauthTokenCommand                         `command:"oauth-token"`
	SSHCode                            SSHCodeCommand                            `command:"ssh-code"`
	AddPluginRepo                      AddPluginRepoCommand                      `command:"add-plugin-repo"`
	RemovePluginRepo                   RemovePluginRepoCommand                   `command:"remove-plugin-repo"`
	ListPluginRepos                    ListPluginReposCommand                    `command:"list-plugin-repos"`
	RepoPlugins                        RepoPluginsCommand                        `command:"repo-plugins"`
	Plugins                            PluginsCommand                            `command:"plugins"`
	InstallPlugin                      InstallPluginCommand                      `command:"install-plugin"`
	UninstallPlugin                    UninstallPluginCommand                    `command:"uninstall-plugin"`
}