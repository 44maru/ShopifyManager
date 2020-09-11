package constants

const INPUT_EXCEL_FILE_PATH = "shopify-input.xlsx"

const FLOW_TYPE_CREATE_INSTANCE = "cancel-order"

//const GET_ORDER_URL_TEMPLATE = "https://%s:%s@penguin-auto-buy-service.myshopify.com/admin/api/2020-07/orders.json?status=any&name=%d"
const GET_ORDER_URL_TEMPLATE = "https://%s:%s@penguin-auto-buy-service.myshopify.com/admin/api/2020-07/orders.json?name=%d"
const CANCEL_ORDER_URL_TEMPLATE = "https://%s:%s@penguin-auto-buy-service.myshopify.com/admin/api/2020-07/orders/%d/cancel.json"
const TRANSACTIONS_URL_TEMPLATE = "https://%s:%s@penguin-auto-buy-service.myshopify.com/admin/api/2020-07/orders/%d/transactions.json"
