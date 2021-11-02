import styled from "styled-components";

export const Container = styled.div`
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  h1 {
    color: #505153;
    margin-bottom: 20px;
  }
`

export const Menu = styled.div`
  width: 800px;
  height: 600px;
  background-color: #fff;

  border-radius: 12px;
  box-shadow: 1px 1px -3px black;

  display: grid;
  grid-template-rows: 1fr 1fr 1fr;
  justify-content: center;
  align-items: center;

  font-size: 12px;

  p {
    margin-top: 50px;
    white-space: pre-line;
    color: #505153
  }
`
