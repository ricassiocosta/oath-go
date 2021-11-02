import React, { useEffect, useState } from "react";
import api from "../../services";
import { Container, Menu } from "./styles";

interface cert {
  id: string,
  serial: string
}

const GetRevogationList = () => {
  const [certs, setCerts] = useState<cert[]>([])

  useEffect(() => {
    async function callApi () {
      const response = await api.get("/crl")
      setCerts(response.data)
    }
    callApi()
  });
  
  
  return (
    <Container>
      <h1>Lista dos Certificados Revogados</h1>
      <Menu>
        <ol>
        {certs.length > 0 ? (
          certs.map(cert => (
            <li key={cert.id}>{cert.serial}</li>
          ))
        ): <></>}
        </ol>
      </Menu>
    </Container>
  )
}

export default GetRevogationList  