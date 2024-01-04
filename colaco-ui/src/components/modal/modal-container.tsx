import styled from "styled-components";
import BaseContainer from "../base-contaner";

const ModalContainer = styled(BaseContainer)`
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: auto;
  height: auto;
  padding: 20px;
  z-index: 1000;
`;

export default ModalContainer;