const baseUrl = 'http://localhost:8090/api/v1'

export async function fetchUtilityBillsByYears(year) {
  const json = await fetch(`${baseUrl}/utility-bills/${year}`).then(response => response.json())

  return json.bills
}
