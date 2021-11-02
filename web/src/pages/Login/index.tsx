import React from "react";

import logo from '../../assets/images/github.svg'
import { LoginPage, GithubLogo, LoginContainer } from "./styles";

const Login = () => {
  return (
    <LoginPage>
      <h1>PKI API</h1>
      <LoginContainer>
      <span>Entrar com o </span>
      <GithubLogo src={logo} alt="testing"/>
      </LoginContainer>
    </LoginPage>
  )
}

export default Login  