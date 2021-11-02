import api from '../'

export const getGithubToken = async (code: any) => {
  const response = await api.get('/callback/github', { params: { code } });
  return response.data.token
}
