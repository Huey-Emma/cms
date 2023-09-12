import requests, json

base_url = "http://127.0.0.1:8080/api"
person_id = 0


class RequestError(Exception):
    def __init__(self, *args: object) -> None:
        super().__init__(*args)


def create_person():
    """Create a person"""
    data = {
            "name": "papa",
            "hobby": "coding"
    } 

    data_json = json.dumps(data)
    resp = requests.post(base_url, data=data_json)

    if resp.status_code // 100 != 2:
        raise RequestError("non-2** response")

    person = resp.json()
    global person_id
    person_id = person["id"]
    print("person created")


def fetch_person():
    """Fetch person by id"""
    resp = requests.get(f'{base_url}/{person_id}')

    if resp.status_code == 404:
        print("not found")
        return 

    if resp.status_code // 100 != 2:
        raise RequestError("non-2** response")

    person = resp.json()
    print(person)


def update_person():
    """Update person"""
    data = {
            "name": "colin",
            "surname": "udoh"
    }

    data_json = json.dumps(data)
    resp = requests.patch(f'{base_url}/{person_id}', data=data_json)

    if resp.status_code // 100 != 2:
        raise RequestError("non-2** response")

    print("person updated")


def delete_person():
    """Delete person"""
    resp = requests.delete(f'{base_url}/{person_id}')

    if resp.status_code == 204:
        print("person deleted")


if __name__ == "__main__":
    print("creating person...")
    create_person()
    print("*" * 20)
    print(f"fetching person by id: {person_id}")
    fetch_person()
    print("*" * 20)
    print("updating person")
    update_person()
    print("*" * 20)
    print("fetching updated person")
    fetch_person()
    print("*" * 20)
    print("deleting person")
    delete_person()
    print("*" * 20)
    print("fetcting deleted person")
    fetch_person()
