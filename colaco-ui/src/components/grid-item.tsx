import styled from "styled-components";

interface GridItem {
  flex?: number | string;
  display?: string;
  alignitems?: string;
  justifycontent?: string;
  padding?: string;
}

const GridItem = styled.div<GridItem>`
  flex: ${(props) => props.flex ?? 1};
  display: ${(props) => props.display ?? "inline-flex"};
  align-items: ${(props) => props.alignitems ?? "center"};
  justify-content: ${(props) => props.justifycontent ?? "center"};
  min-height: 0;
  min-width: 0;
  padding: ${(props) => props.padding ?? "0"};
`;

export default GridItem;
