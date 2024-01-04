interface Image {
    url: string;
    tagline: string;
    productName?: string;
}

const IMAGES: { [key: string]: Image } = {
    FIZZ: {
        url: new URL('./FIZZ.png', import.meta.url).href,
        tagline: "High speed, low drag",
        productName: "Fizz",
    },
    POP: {
        url: new URL('./POP.png', import.meta.url).href,
        tagline: "Tried and true",
        productName: "Pop",
    },
    COLA: {
        url: new URL('./COLA.png', import.meta.url).href,
        tagline: "The one you know and love",
        productName: "Cola",
    },
    MEGAPOP: {
        url: new URL('./MEGAPOP.png', import.meta.url).href,
        tagline: "I hope you have health insurance",
        productName: "Mega Pop",
    },
};

export default IMAGES;