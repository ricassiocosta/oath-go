import React from "react";
import { Container, Menu } from "./styles";

const SeeCertificate = () => {
  const certificate = localStorage.getItem("certificate")
  
  return (
    <Container>
      <h1>Certificado:</h1>
      <Menu>
        <p>{certificate}</p>
      </Menu>
    </Container>
  )
}

export default SeeCertificate  