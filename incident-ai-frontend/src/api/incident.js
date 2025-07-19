import axios from "axios";

export const BASE_URL = "http://localhost:5174/incidents";

export const createIncident = (data) => axios.post(BASE_URL, data);

export const fetchIncidents = () => axios.get(BASE_URL);

export const getIncidentById = (id) => axios.get(`${BASE_URL}/${id}`);
