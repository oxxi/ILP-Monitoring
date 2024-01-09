export const getAll = async () => {
  return fetch(
    'https://c2cjpykaocnu3e74ztt6bvcnam0gxfbx.lambda-url.us-east-1.on.aws/'
  )
    .then((resp) => {
      if (!resp.ok) {
        throw new Error('Error ');
      }
      return resp.json();
    })
    .catch((err) => {
      throw new Error(err.message);
    });
};
