# gitbook-integration

test sync

test from gitbook



## Create a new user

<mark style="color:green;">`POST`</mark> `/users`

\<Description of the endpoint>

**Headers**

| Name          | Value              |
| ------------- | ------------------ |
| Content-Type  | `application/json` |
| Authorization | `Bearer <token>`   |

**Body**

| Name   | Type   | Description      |
| ------ | ------ | ---------------- |
| `name` | string | Name of the user |
| `age`  | number | Age of the user  |

**Response**

{% tabs %}
{% tab title="200" %}
```json
{
  "id": 1,
  "name": "John",
  "age": 30
}
```
{% endtab %}

{% tab title="400" %}
```json
{
  "error": "Invalid request"
}
```
{% endtab %}
{% endtabs %}

{% openapi src="petstore.yaml" path="/store/inventory" method="get" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}

{% openapi src="petstore.yaml" path="/store/order" method="post" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}

{% openapi src="petstore.yaml" path="/store/order/{orderId}" method="get" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}

{% openapi src="petstore.yaml" path="/store/order/{orderId}" method="delete" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}

{% openapi src="petstore.yaml" path="/pet" method="PUT" expanded="true" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}

{% openapi src="petstore.yaml" path="/pet" method="POST" expanded="true" %}
[petstore.yaml](petstore.yaml)
{% endopenapi %}
