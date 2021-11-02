import React, { useEffect } from "react";
import { useHistory, useLocation } from "react-router-dom";
import { authenticate } from "../../services/auth";
import { getGithubToken } from "../../services/github";

const Callback = () => {
  const query = new URLSearchParams(useLocation().search)
  const history = useHistory()

  useEffect(() => {
    async function callApi () {
      const githubToken = await getGithubToken(query.get('code'))
      const response  = await authenticate(githubToken)
      
      localStorage.setItem("token", response.token)
      
      history.replace('/')
    }
    callApi()
  });

  return (
    <></>
  )
}

export default Callback
