import styled from 'styled-components';

interface ContentCardProps {
    image: string;
    children?: React.ReactNode;
    onClick?: () => void;
}

const StyledContentCard = styled.div<ContentCardProps>`
    background-image: url(${props => props.image});
    background-size: cover;
    background-position: center;
    padding: 5px;
    height: calc(100% - 30px); // subtracting the padding
    width: calc(100% - 30px); // subtracting the padding
    border-radius: 20px;

    &:hover {
        cursor: pointer;
        box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.75);
    }
`;

const ContentCard = ({ image, children, onClick }: ContentCardProps) => (
    <StyledContentCard onClick={onClick} image={image}>
        {children}
    </StyledContentCard>
);

export default ContentCard;