import styled from "styled-components";

const DetailItem = styled.details`
    width: 100%;
    height: 100%;
    border: 1px solid ${props => props.theme.colors.secondary};
    border-radius: 10;
    overflow: hidden;
    text-align: left;

    summary {
        background-color: ${props => props.theme.colors.primary};
        color: ${props => props.theme.colors.secondary};
        border: none;
        border-radius: 0;
    }
`;

export default DetailItem;