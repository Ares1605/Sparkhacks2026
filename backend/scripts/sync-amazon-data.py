from amazonorders.session import AmazonSession
from amazonorders.orders import AmazonOrders
from dotenv import load_dotenv
import os
import json
import sys

load_dotenv()

# We would like to mock the Amazon data, for testing purposes
if os.getenv("AMAZON_MOCK_SYNC") == "1":
    script_directory = os.path.dirname(os.path.abspath(sys.argv[0]))
    print(script_directory)
    with open(os.path.join(script_directory, "amazon-mock-data.json"), "r") as mock_file:
        print(mock_file.read(), flush=True)
        exit(0)

amazon_session = AmazonSession(os.getenv("AMAZON_USERNAME"),
                               os.getenv("AMAZON_PASSWORD"),
                               otp_secret_key=os.getenv("AMAZON_OTP_KEY"))
amazon_session.login()

amazon_orders = AmazonOrders(amazon_session)

orders = amazon_orders.get_order_history(
    full_details=True,
)

serialized_orders = []

for order in orders:
    serialized = {
        "amazon_discount": order.amazon_discount,
        "order_number": order.order_number,
        "coupon_savings": order.coupon_savings,
        "free_shipping": order.free_shipping,
        "gift_card": order.gift_card,
        "shipping_total": order.shipping_total,
        "subtotal": order.subtotal,
        "order_placed_date": order.order_placed_date.isoformat(),
        "order_details_link": order.order_details_link,
        "items": []
    }
    for item in order.items:
        serialized["items"].append({
            "image_link": item.image_link,
            "price": item.price,
            "quantity": item.quantity,
            "title": item.title,
            "condition": item.condition,
            "link": item.link
        })

    serialized_orders.append(serialized)

print(json.dumps(serialized_orders), flush=True)
