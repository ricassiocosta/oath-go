import React from "react";

import logo from '../../assets/images/github.svg'
import { LoginPage, GithubLogo, LoginContainer } from "./styles";
import { constants } from "../../constants";
import { env } from "../../env";

const Login = () => {
  return (
    <LoginPage>
      <h1>PKI API</h1>
      <LoginContainer href={`${constants.AUTHORIZE_URL}?client_id=${env.GITHUB_CLIENT_ID}`}>
      <span>Entrar com o </span>
      <GithubLogo src={logo} alt="testing"/>
      </LoginContainer>
    </LoginPage> 
  )
}

export default Login  