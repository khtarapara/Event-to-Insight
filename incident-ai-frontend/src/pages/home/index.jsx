import {
  Container,
  Typography,
  Box,
  Divider,
  CircularProgress,
  Alert,
  Button,
} from "@mui/material";
import IncidentForm from "./IncidentForm";
import IncidentTable from "./IncidentTable";
import { useCallback, useEffect, useState } from "react";
import { fetchIncidents } from "../../api/incident";

export default function Home() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [showCreateForm, setShowCreateForm] = useState(false);

  const loadIncidents = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const res = await fetchIncidents();
      setData(res.data);
    } catch (err) {
      console.error("Error fetching incidents:", err);
      setError("Failed to load incidents.");
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    loadIncidents();
  }, [loadIncidents]);

  return (
    <Container maxWidth="md">
      <Typography variant="h4" gutterBottom>
        Incident Dashboard
      </Typography>
      {showCreateForm ? (
        <IncidentForm
          onIncidentCreated={loadIncidents}
          onCancel={() => setShowCreateForm(false)}
        />
      ) : (
        <Button onClick={() => setShowCreateForm(true)} title="create-incident">
          Create Incident
        </Button>
      )}
      <Box my={4}>
        <Divider />
      </Box>
      {loading && <CircularProgress />}
      {error && <Alert severity="error">{error}</Alert>}
      {!loading && !error && <IncidentTable data={data} />}
    </Container>
  );
}
