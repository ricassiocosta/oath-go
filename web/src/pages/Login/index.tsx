import React from "react";

import logo from '../../assets/images/github.svg'
import { LoginPage } from "./styles";

const Login = () => {
  return (
    <LoginPage>
      <p>Testing</p>
      <img src={logo} alt="testing"/>
    </LoginPage>
  )
}

export default Login