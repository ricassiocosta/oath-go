import React, { useEffect } from "react";
import { BrowserRouter, Redirect, Route, Switch } from "react-router-dom";
import AddToCRL from "./pages/addToCRL";
import Callback from "./pages/Callback";
import Dashboard from "./pages/Dashboard";
import Generate from "./pages/GenerateCertificate";
import Login from "./pages/Login";
import GetRevogationList from "./pages/RevogationList";
import SeeCertificate from "./pages/SeeCertificate";

const AuthenticatedRoute = ({ component: Component, ...rest }: any) => {
  const isAuthenticated = localStorage.getItem("token")

  useEffect(() => {
  }, [isAuthenticated]);

  return (
    <Route
    {...rest}
    render={(props) => isAuthenticated ? (
            <Component {...props} {...rest} />
        ) : (
            <Redirect to={'/login'}/>
        )
    }
    />
  );
}


const Routes = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/login" component={Login}/>
        <Route path="/callback" component={Callback}/>
        <AuthenticatedRoute path="/" exact component={Dashboard} />
        <AuthenticatedRoute path="/revogar" exact component={AddToCRL} />
        <AuthenticatedRoute path="/obterrevogados" exact component={GetRevogationList} />
        <AuthenticatedRoute path="/gerar" exact component={Generate} />
        <AuthenticatedRoute path="/certificate" exact component={SeeCertificate} />
      </Switch>
    </BrowserRouter>
  )
}

export default Routes