package skip_messages

const SkipAppsMessage = `Skipping this test because config.IncludeApps is set to 'false'.`
const SkipBackendCompatibilityMessage = `Skipping this test because config.IncludeBackendCompatibility is set to 'false'.`
const SkipContainerNetworkingMessage = `Skipping this test because Config.IncludeContainerNetworking is set to 'false'.`
const SkipDetectMessage = `Skipping this test because config.IncludeDetect is set to 'false'.`
const SkipDockerMessage = `Skipping this test because config.IncludeDocker is set to 'false'.
NOTE: Ensure Docker containers are enabled on your platform before enabling this test.`
const SkipInternetDependentMessage = `Skipping this test because config.IncludeInternetDependent is set to 'false'.
NOTE: Ensure that your platform has access to the internet before running this test.`
const SkipInternetlessMessage = `Skipping this test because config.IncludeInternetless is set to 'false'.
NOTE: Ensure that your platform does not have access to the internet before running this test.`
const SkipPrivateDockerRegistryMessage = `Skipping this test because config.IncludePrivateDockerRegistry is set to 'false'.
NOTE: Ensure that you've provided values for config.PrivateDockerRegistryImage, config.PrivateDockerRegistryUsername,
and config.PrivateDockerRegistryPassword before running this test.`
const SkipCredhubMessage = `Skipping this test because Config.CredhubMode is not set to either 'assisted' or 'non-assisted'.
NOTE: Ensure instance identity credential is turned on and CredHub is deployed before enabling this test`
const SkipAssistedCredhubMessage = `Skipping this test because Config.CredhubMode is not set to 'assisted'.
NOTE: Ensure instance identity credential is turned on and CredHub is deployed before enabling this test`
const SkipNonAssistedCredhubMessage = `Skipping this test because Config.CredhubMode is not set to 'non-assisted'.
NOTE: Ensure instance identity credential is turned on and CredHub is deployed before enabling this test`
const SkipRouteServicesMessage = `Skipping this test because config.IncludeRouteServices is set to 'false'.
NOTE: Ensure that route services are enabled on your platform before running this test.`
const SkipRoutingMessage = `Skipping this test because config.IncludeRouting is set to 'false'.`
const SkipTCPRoutingMessage = `Skipping this test because Config.IncludeTCPRouting is set to 'false'.`
const SkipSecurityGroupsMessage = `Skipping this test because config.IncludeSecurityGroups is set to 'false'.
NOTE: Ensure that your platform restricts internal network traffic by default in order to run this test.`
const SkipServicesMessage = `Skipping this test because config.IncludeServices is set to 'false'.`
const SkipSSHMessage = `Skipping this test because config.IncludeSsh is set to 'false'.
NOTE: Ensure that your platform is deployed with a Diego SSH proxy in order to run this test.`
const SkipSSOMessage = `Skipping this test because config.IncludeSSO is not set to 'true'.
NOTE: Ensure that your platform is running UAA with SSO enabled before enabling this test.`
const SkipTasksMessage = `Skipping this test because config.IncludeTasks is set to 'false'.
NOTE: Ensure tasks are enabled on your platform before enabling this test.`
const SkipV3Message = `Skipping this test because config.IncludeV3 is set to 'false'.
NOTE: Ensure that the v3 api features are enabled on your platform before running this test.`
const SkipWindowsMessage = `Skipping this test because config.IncludeWindows is set to 'false'.
NOTE: Ensure that your deployment includes at least one Windows cell before enabling this test.`
const SkipWindowsContextPathsMessage = `Skipping this test because config.UseWindowsContextPath is set to 'false'.
NOTE: Ensure that your deployment includes at least one Windows cell before enabling this test.`
const SkipIsolationSegmentsMessage = `Skipping this test because Config.IncludeIsolationSegments is set to 'false'`
const SkipRoutingIsolationSegmentsMessage = `Skipping this test because Config.IncludeRoutingIsolationSegments is set to 'false'.`
const SkipZipkinMessage = `Skipping this test because Config.IncludeZipkin is set to 'false'`
const SkipServiceDiscoveryMessage = `Skipping this test because Config.IncludeServiceDiscovery is set to 'false'.`
const SkipServiceInstanceSharingMessage = `Skipping this test because Config.IncludeServiceInstanceSharing is set to 'false'.`
const SkipCapiExperimentalMessage = `Skipping this test because Config.IncludeCapiExperimental is set to 'false'.`
const SkipCapiNoBridgeMessage = `Skipping this test because Config.IncludeCapiNoBridge is set to 'false'.`
const SkipSSHOnWindows2012R2Message = `cf ssh does not work on windows2012R2`
const SkipWindowsTasksMessage = `Skipping Windows tasks tests (requires diego-release v1.20.0 and above)`
const SkipNoAlternateStacksMessage = `Skipping this test because no 'alternate_stacks' are specified.`
