import styled from "styled-components";

interface ButtonProps {
  backgroundColor?: string;
  borderColor?: string;
  selected?: boolean;
  rounded?: boolean;
  roundedLeft?: boolean;
  roundedRight?: boolean;
}

const Button = styled.button<ButtonProps>`
  display: inline-flex;
  align-items: center;
  justify-content: center;

  background-color: ${props => props.backgroundColor ?? "transparent"};
  color: ${props => props.theme.colors.secondary};
  height: 100%;
  width: auto;
  border-radius: ${props => {
    if (props.roundedLeft) {
      return "5px 0 0 5px";
    }

    if (props.roundedRight) {
      return "0 5px 5px 0";
    }

    if (props.rounded) {
      return "5px";
    }
  }};
  border: 1px solid ${props => props.borderColor ?? "transparent"};
  padding: 0.6em 1.2em;
  cursor: pointer;
  overflow: hidden;

  &:hover {
    background-color: ${props => props.theme.colors.secondary};
    color: ${props => props.theme.colors.primary};
  }

  ${props => props.selected && `
    background-color: ${props.theme.colors.secondary};
    color: ${props.theme.colors.primary};
  `}

  img {
    height: 100%;
    width: 100%;
    object-fit: contain;
  }
`;

export default Button;