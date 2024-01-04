import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import BaseContainer from "../base-contaner";
import Button from "../button";
import GridContainer from "../grid-container";
import GridItem from "../grid-item";
import { faShoppingCart } from "@fortawesome/free-solid-svg-icons";
import DropdownButton, { DropdownItem } from "../dropdown-button";
import { useState } from "react";
import TextInput from "../text-input";
import Modal from "../modal/modal";
import { User } from "../../queries/get-user";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import purchaseCart from "../../queries/purchase-cart";
import addFunds from "../../queries/add-funds";

interface ToolbarProps {
    cartItems?: DropdownItem[];
    user: User;
    setCartItems: (arg0: DropdownItem[]) => void;
}

const Toolbar = ({ cartItems, user, setCartItems }: ToolbarProps) => {
    const queryClient = useQueryClient();

    const [showAddFundsModal, setShowAddFundsModal] = useState(false);
    const [newFunds, setNewFunds] = useState(0);

    const purchaseCartMutation = useMutation({
        mutationFn: (items: DropdownItem[]) => {
            queryClient.invalidateQueries({
                queryKey: ['getUserAccount', user.id]
            });
            setCartItems([]);
            return purchaseCart(user.id, items);
        },
    });

    const addFundsMutation = useMutation({
        mutationFn: (amount: number) => {
            queryClient.invalidateQueries({
                queryKey: ['getUserAccount', user.id]
            });
            return addFunds(user.id, amount);
        },
    });

    return (
        <BaseContainer>
            <GridContainer>
                <GridItem justifycontent="left">
                    <GridContainer layout="column">
                        <GridItem>
                            <h1>ColaCo</h1>
                        </GridItem>
                        <GridItem>
                            <p>Elevate your taste</p>
                        </GridItem>
                    </GridContainer>
                </GridItem>
                <GridItem justifycontent="flex-start" flex={4}>
                    <GridContainer layout="column">
                        <GridItem>
                            <h3>Account Balance</h3>
                        </GridItem>
                        <GridItem>
                            <h3>$</h3><TextInput value={user.balance} defaultValue={"1.00"} type="text" disabled />
                        </GridItem>
                    </GridContainer>
                </GridItem>
                <GridItem justifycontent="flex-end">
                    <DropdownButton
                        backgroundColor="#D58384"
                        items={cartItems ?? []}
                        onSelectCheckout={(items: DropdownItem[]) => {
                            purchaseCartMutation.mutate(items);
                        }}
                    >
                        <FontAwesomeIcon size="2x" icon={faShoppingCart} />
                    </DropdownButton>
                    <Button onClick={() => setShowAddFundsModal(!showAddFundsModal)} backgroundColor="#EE933D">
                        <GridContainer layout="column">
                            <GridItem>
                                <h3>Add</h3>
                            </GridItem>
                            <GridItem>
                                <h3>Funds</h3>
                            </GridItem>
                        </GridContainer>
                    </Button>
                    <Modal showModal={showAddFundsModal} setShowModal={setShowAddFundsModal}>
                        <GridContainer layout="column">
                            <GridItem>
                                <h3>Add New Balance:</h3>
                            </GridItem>
                            <GridItem>
                                <h3>$</h3><TextInput type="text" defaultValue={"1"} value={newFunds} onChange={(e) => {
                                    setNewFunds(Number(e.target.value));
                                }
                                } />
                            </GridItem>
                            <GridItem>
                                <Button backgroundColor="#10AC84" rounded onClick={() => addFundsMutation.mutate(newFunds, {
                                    onSuccess: () => {
                                        setShowAddFundsModal(false);
                                    }
                                })}>
                                    <h4>
                                        Submit
                                    </h4>
                                </Button>
                            </GridItem>
                        </GridContainer>
                    </Modal>
                </GridItem>
            </GridContainer>
        </BaseContainer>
    );
};

export default Toolbar;
