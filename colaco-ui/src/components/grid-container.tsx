import styled from "styled-components";

interface GridContainerProps {
  layout?: "row" | "column";
  orientation?: "flex-start" | "flex-end" | "center" | "space-between" | "space-around";
}

const GridContainer = styled.div<GridContainerProps>`
  display: flex;
  height: 100%;
  width: 100%;
  flex-direction: ${(props) => props.layout ?? "row"};
  justify-content: ${(props) => props.orientation ?? "flex-start"};

  h1, h2, h3 {
    font-family: "Changa One", sans-serif;
  }

  h4, h5, h6, p, span {
    font-family: "Pacifico", cursive;
  }
}`;

export default GridContainer;
