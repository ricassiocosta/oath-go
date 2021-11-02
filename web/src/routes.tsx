import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Callback from "./pages/Callback";
import Login from "./pages/Login";

const Routes = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/login" component={Login}/>
        <Route path="/callback" component={Callback}/>
      </Switch>
    </BrowserRouter>
  )
}

export default Routes