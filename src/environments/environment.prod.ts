const redirect_uri = window.location.protocol + '//' + window.location.host + '/clusters';
const oauth = window.location.protocol + '//' + window.location.host + '/dex/auth';
const scope: string[] = ['openid', 'email', 'profile', 'groups'];

export const environment = {
  name: 'prod',
  production: true,
  restRoot : '/api/v1',
  restRootV3 : '/api/v3',
  digitalOceanRestRoot : 'https://api.digitalocean.com/v2',
  coreOSdexAuth : oauth + '?response_type=id_token&client_id=kubermatic&redirect_uri='
  + redirect_uri + '&scope=' + scope.join(' ') + '&nonce=random',
};
