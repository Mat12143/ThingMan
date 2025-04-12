async function update(item) {
  let resp = null;

  try {
    resp = await fetch("/update/" + item.id, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });
  } catch (e) {
    console.error(e);
    return null;
  }

  if (!resp.ok) {
    console.error(resp);
    return null;
  }
  try {
    // When the item is updated
    const json = await resp.json();
    return json;
  } catch (e) {
    // If the item is deleted
    return {};
  }
}

async function addNew(id) {
  const form = document.forms["form-" + id];
  const name = form.name.value;
  const location = id;
  const quantity = form.quantity.value;

  data = {
    name: name,
    location: Number(location),
    quantity: Number(quantity),
  };

  try {
    const resp = await fetch("/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (resp.ok) {
      form.reset();
      return await resp.json();
    }
  } catch (e) {
    console.error("Cannot add a new item. Try to refresh");
    return nil;
  }
}

async function handleMove(item, position) {
  console.log(item, position);
}
