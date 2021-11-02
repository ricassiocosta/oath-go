import React, { useState } from "react";
import { useHistory } from "react-router";
import api from "../../services";
import { Container, Menu, Option } from "./styles";

const Generate = () => {
  const [csr, setCSR] = useState('')
  const history = useHistory()

  async function handleSubmit() {
    const token = localStorage.getItem("token")
    if (token == null) {
      history.replace("/login")
    }

    const data = {
      csr: csr
    }

    const response = await api.post("/signcertificate", data, {
      headers: {
        Authorization: `Bearer ${token}`,
      }
    })

    console.log(response)

    localStorage.setItem("certificate", response.data)

    history.replace("/certificate")
  }
  
  return (
    <Container>
      <h1>Gerar Certificado</h1>
      <Menu>
        <p>Digite o conteúdo da CSR:</p>
        <input
            placeholder="Conteúdo da CSR"
            value={csr}
            onChange={e => setCSR(e.target.value)}
          />
        <Option onClick={handleSubmit}>Enviar</Option>
      </Menu>
    </Container>
  )
}

export default Generate  