const generateID = (len = 6) => {
  var id = "";

  for(let i = 0; i < len; i++) {
    id += String(Math.floor(Math.random() * 10));
  }

  return id;
}

export { generateID };
