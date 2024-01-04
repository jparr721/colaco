import React, { useEffect, useRef, useState } from 'react';
import styled from 'styled-components';
import BaseContainer from './base-contaner';
import Button from './button';

const DropdownMenu = styled(BaseContainer)`
  display: block;
  position: absolute;
  width: auto;
  height: auto;
  background-color: ${props => props.theme.colors.primary};
  color: ${props => props.theme.colors.secondary};
  z-index: 1;
`;

export interface DropdownItem {
    id: string;
    product_name: string;
    price: number;
}

const DropdownItem = styled.div<DropdownItem>`
  padding: 0.6em 1.2em;
  cursor: pointer;

  &:hover {
    background-color: ${props => props.theme.colors.secondary};
    color: ${props => props.theme.colors.primary};
  }
`;

interface DropdownButtonProps {
    items: DropdownItem[];
    onSelectCheckout: (item: DropdownItem[]) => void;
    children: React.ReactNode;
    backgroundColor?: string;
}

const DropdownButton = ({ items, onSelectCheckout, children, backgroundColor }: DropdownButtonProps) => {
    const [isOpen, setIsOpen] = useState(false);
    const buttonRef = useRef(null);
    const dropdownRef = useRef(null);
    const [menuPosition, setMenuPosition] = useState({ top: 0, left: 0 });

    const toggleDropdown = () => {
        if (buttonRef.current) {
            // Get the button position
            const buttonRect = buttonRef.current.getBoundingClientRect();

            // ... and then set the menu position to the bottom of the button.
            // Is this somewhat goofy? Yeah, but it works quite well honestly.
            setMenuPosition({
                top: buttonRect.bottom + window.scrollY,
                left: items.length > 0 ? buttonRect.left + window.scrollX - 200 : buttonRect.left + window.scrollX
            });
        }
        setIsOpen(!isOpen);
    };

    const handleItemSelect = (item: string) => {
        // onSelectItem(item);

        const found = items.findIndex((i) => i.product_name === item);

        if (found !== -1) {
            items.splice(found, 1);
        }

        setIsOpen(false);
    };

    const handleCheckout = () => {
        onSelectCheckout(items);
        setIsOpen(false);
    };

    // If the user clicks outside of the dropdown, close it
    useEffect(() => {
        const handleOutsideClick = (event: MouseEvent) => {
            if (buttonRef.current && !buttonRef.current.contains(event.target) && dropdownRef.current && !dropdownRef.current.contains(event.target)) {
                setIsOpen(false);
            }

        };
        document.addEventListener('mousedown', handleOutsideClick);
        return () => document.removeEventListener('mousedown', handleOutsideClick);
    }, [buttonRef, dropdownRef]);

    const frequencyMap = items.reduce((acc, item) => {
        acc[item.product_name] = (acc[item.product_name] || 0) + 1;
        return acc;
    }, {} as Record<string, number>);


    return (
        <>
            <Button backgroundColor={backgroundColor} ref={buttonRef} onClick={() => toggleDropdown()}>
                {children}
            </Button>
            {
                isOpen &&
                <DropdownMenu style={{ top: `${menuPosition.top}px`, left: `${menuPosition.left}px` }} ref={dropdownRef}>
                    {
                        items.length > 0 ?
                            Object.entries(frequencyMap).map(([productName, count], index) => (
                                <DropdownItem key={index} onClick={() => handleItemSelect(productName)}>
                                    <h4>{productName} {count > 1 ? `x${count}` : 'x1'} | ${
                                        Number(items.find(i => i.product_name == productName)!.price) * count}</h4>
                                </DropdownItem>

                            )) : <h4>Cart Is Empty.</h4>}
                    <Button backgroundColor='#10AC84' onClick={handleCheckout}>
                        <h3>Checkout</h3>
                    </Button>
                    {
                        items.length > 0 &&
                        <Button backgroundColor='#D58384' onClick={() => items.length = 0}>
                            <h3>Clear</h3>
                        </Button>
                    }
                </DropdownMenu>
            }
        </>
    );
};

export default DropdownButton;
