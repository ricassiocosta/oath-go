import React, { useState } from "react";
import { useHistory } from "react-router";
import api from "../../services";
import { Container, Menu, Option } from "./styles";

const AddToCRL = () => {
  const [serial, setSerial] = useState('')
  const history = useHistory()

  async function handleSubmit() {
    const token = localStorage.getItem("token")
    if (token == null) {
      history.replace("/login")
    }

    const data = {
      serial: serial
    }

    await api.post("/crl", data, {
      headers: {
        Authorization: `Bearer ${token}`,
      }
    })

    history.replace("/")
  }
  
  return (
    <Container>
      <h1>Revogar Certificado</h1>
      <Menu>
        <p>Digite o serial do certificado que deseja revogar:</p>
        <input
            placeholder="Serial do certificado"
            value={serial}
            onChange={e => setSerial(e.target.value)}
          />
        <Option onClick={handleSubmit}>Enviar</Option>
      </Menu>
    </Container>
  )
}

export default AddToCRL  