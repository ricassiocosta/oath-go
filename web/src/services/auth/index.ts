import api from '../'

export const authenticate = async (githubToken: any) => {
  const response = await api.get('/auth', { params: { githubToken } });
  return response.data
}