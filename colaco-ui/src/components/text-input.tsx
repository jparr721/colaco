import styled from "styled-components";

const TextInput = styled.input`
    background-color: ${props => props.theme.colors.accentLight};
    color: ${props => props.theme.colors.secondary};
    height: 80%;
    width: auto;
    border: 0px solid ${props => props.theme.colors.secondary};
    border-radius: 50px;
    padding: 0.6em 1.2em;
    font-size: 1.2em;
`;

export default TextInput;