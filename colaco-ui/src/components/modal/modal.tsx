import React from "react";
import ModalBackdrop from "./modal-backdrop";
import ModalContainer from "./modal-container";

interface ModalProps {
  children: React.ReactNode;
  showModal: boolean;
  setShowModal: (arg0: boolean) => void;
}

const Modal = ({ children, showModal, setShowModal }: ModalProps) => {
  return (
    <>
      {showModal && (
        <>
          <ModalBackdrop onClick={() => setShowModal(false)} />
          <ModalContainer>{children}</ModalContainer>
        </>
      )}
    </>
  );
};

export default Modal;
