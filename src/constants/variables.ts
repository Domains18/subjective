import dotenv from "dotenv";

dotenv.config();

const whatsappToken = process.env.WHATSAPP_ACCESS_TOKEN;
const metaVersion = process.env.META_VERSION;
const phoneNumberID = process.env.PHONE_NUMBER_ID;
const port = process.env.PORT;

export { whatsappToken, metaVersion, phoneNumberID, port };
