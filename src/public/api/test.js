export default async function handler(request, response) {

  const data = { message: "Hello from Vercel API!" };

  response.status(200).json(data);

}
