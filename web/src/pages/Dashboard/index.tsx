import React from "react"
import { Link } from "react-router-dom"

import { DashboardContainer, Menu, Option } from "./styles"

const Dashboard = () => {
  return (
    <DashboardContainer>
      <h1>Dashboard</h1>
      <Menu>
        <Link to="/gerar">
          <Option>Gerar Certificado</Option>
        </Link>
        <Link to="/revogar">
          <Option>Revogar Certificado</Option>
        </Link>
        <Link to="/obterrevogados">
          <Option>Obter Lista de Certificados Revogados</Option>
        </Link>
      </Menu>
    </DashboardContainer>
  )
}

export default Dashboard  