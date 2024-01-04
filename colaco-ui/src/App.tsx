import './App.css';
import BaseContainer from './components/base-contaner';
import ContentCard from './components/content-card';
import GridContainer from './components/grid-container';
import GridItem from './components/grid-item';
import IMAGES from "../public/images";
import Toolbar from './components/toolbar/toolbar';
import Button from './components/button';
import { CookiesProvider, useCookies } from 'react-cookie';
import { useQuery } from '@tanstack/react-query';
import createUser from './queries/create-user';
import getUser from './queries/get-user';
import getAllSodas, { Soda } from './queries/get-all-sodas';
import { useState } from 'react';
import Modal from './components/modal/modal';
import { DropdownItem } from './components/dropdown-button';

function App() {
  const [cookies, setCookie] = useCookies(['user-account']);
  const [productDetailsModalOpen, setProductDetailsModalOpen] = useState(false);
  const [selectedProduct, setSelectedProduct] = useState<Soda>({
    id: "",
    product_name: "",
    description: "",
    cost: 0,
    current_quantity: 0,
    max_quantity: 0,
    created_at: "",
    updated_at: "",
  });
  const [cartItems, setCartItems] = useState<DropdownItem[]>([]);

  const selectProduct = (key: string) => {
    const image = IMAGES[key as keyof typeof IMAGES];
    const dbData = getAllSodasQuery.data?.find((soda: Soda) => soda.product_name === image.productName);

    setSelectedProduct(dbData!);
    setProductDetailsModalOpen(!productDetailsModalOpen);
  };

  // Load a new cookie on page load
  const createUserAccountQuery = useQuery({
    queryKey: ['createUserAccount'],
    queryFn: createUser,
    enabled: cookies['user-account'] === undefined,
  });

  if (createUserAccountQuery.data) {
    setCookie('user-account', createUserAccountQuery.data.id, { path: '/' });
  }

  // Load relevant user details
  const getUserAccountQuery = useQuery({
    queryKey: ['getUserAccount', cookies['user-account']],
    queryFn: getUser,
    refetchInterval: 1000,
    enabled: cookies['user-account'] !== undefined,
  });

  // Load the sodas
  const getAllSodasQuery = useQuery({
    queryKey: ['getAllSodas'],
    queryFn: getAllSodas,
    refetchInterval: 1000,
  });

  const addProductToCart = (dbData: Soda) => {
    setCartItems([...cartItems, { id: dbData?.id, product_name: dbData?.product_name, price: dbData?.cost }]);
  };

  return (
    <CookiesProvider>
      <BaseContainer>
        <GridContainer layout='column'>
          <GridItem flex={1}>
            {(getUserAccountQuery.isLoading || createUserAccountQuery.isLoading) ? <p>Loading</p> : <Toolbar user={getUserAccountQuery.data!} cartItems={cartItems} setCartItems={setCartItems} />}
          </GridItem>
          <GridItem flex={10} padding="5px">
            <GridContainer layout="row">
              {
                getAllSodasQuery.isLoading ? <p>Loading</p> :
                  Object.keys(IMAGES).map((key) => {
                    const image = IMAGES[key as keyof typeof IMAGES];
                    const dbData = getAllSodasQuery.data?.find((soda: Soda) => soda.product_name === image.productName);

                    return (
                      <GridItem>
                        <GridContainer layout="column">
                          <GridItem flex={4}>
                            <ContentCard onClick={() => selectProduct(key)} image={image.url} />
                          </GridItem>
                          <Modal showModal={productDetailsModalOpen} setShowModal={setProductDetailsModalOpen}>
                            <GridContainer layout="column">
                              <GridItem>
                                <h3>{selectedProduct.product_name}:</h3>
                              </GridItem>
                              <GridItem>
                                <p>{selectedProduct.description}</p>
                              </GridItem>
                              <GridItem>
                                <Button onClick={() => setProductDetailsModalOpen(!productDetailsModalOpen)}>
                                  Close
                                </Button>
                              </GridItem>
                            </GridContainer>
                          </Modal>
                          <GridItem>
                            <GridContainer layout="column">
                              <GridItem>
                                <h1>{image.productName}</h1>
                              </GridItem>
                              <GridItem>
                                <h4>{image.tagline}</h4>
                              </GridItem>
                              <GridItem>
                                <Button borderColor='black' roundedLeft onClick={() => addProductToCart(dbData!)}>
                                  <h4>{dbData?.current_quantity}</h4>
                                </Button>
                                <Button borderColor='black' roundedRight>
                                  <h4>${dbData?.cost}</h4>
                                </Button>
                              </GridItem>
                            </GridContainer>
                          </GridItem>
                        </GridContainer>
                      </GridItem>
                    );
                  })
              }
            </GridContainer>
          </GridItem>
        </GridContainer>
      </BaseContainer>
    </CookiesProvider>
  );
}

export default App;
