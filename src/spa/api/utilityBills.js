const baseUrl = import.meta.env.API_BASE_URL

export async function fetchUtilityBillsByYears(year) {
  const json = await fetch(`${baseUrl}/utility-bills/${year}`).then(response => response.json())

  return json.bills
}
