import styled from "styled-components";

export const LoginPage = styled.div`
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

export const LoginContainer = styled.button`
  border: none;
  border-radius: 12px;
  font-size: 18px;
  background-color: #505153;

  display: flex;
  justify-content: space-evenly;
  align-items: center;
  width: 180px;
  height: 60px;

  &:hover {
    filter: brightness(0.8);
    cursor: pointer;
  }

  span {
    color: white;
  }

`

export const GithubLogo = styled.img`
  width: 32px;
  height: 32px;
`
