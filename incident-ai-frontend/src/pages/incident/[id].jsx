import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getIncidentById } from "../../api/incident";
import {
  Container,
  Typography,
  CircularProgress,
  Alert,
  Box,
  Paper,
} from "@mui/material";

export default function IncidentDetail() {
  const { id } = useParams();
  const [incident, setIncident] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchIncident = async () => {
      setLoading(true);
      setError(null);
      try {
        const res = await getIncidentById(id);
        setIncident(res.data);
      } catch {
        setError("Failed to load incident.");
      } finally {
        setLoading(false);
      }
    };
    fetchIncident();
  }, [id]);

  if (loading) return <CircularProgress />;
  if (error) return <Alert severity="error">{error}</Alert>;
  if (!incident) return null;

  return (
    <Container maxWidth="sm">
      <Typography variant="h5" gutterBottom>
        Incident #{incident.id}
      </Typography>
      <Paper
        sx={{
          p: 3,
          borderRadius: 2,
          boxShadow: 3,
          backgroundColor: "#fafafa",
          fontFamily: "'Inter', sans-serif",
        }}
      >
        <Box mb={2}>
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>Title</strong>
          </Typography>
          <Typography data-testid="incident-title" variant="body1" paragraph>
            {incident.title}
          </Typography>
        </Box>
        <Box mb={2}>
          <strong>Description:</strong> {incident.description}
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>Description</strong>
          </Typography>
          <Typography
            data-testid="incident-description"
            variant="body1"
            paragraph
          >
            {incident.description}
          </Typography>
        </Box>
        <Box mb={2}>
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>Affected Service</strong>
          </Typography>
          <Typography
            data-testid="incident-affected-service"
            variant="body1"
            paragraph
          >
            {incident.affected_service}
          </Typography>
        </Box>
        <Box mb={2}>
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>AI Severity</strong>
          </Typography>
          <Typography
            data-testid="incident-ai-severity"
            variant="body1"
            paragraph
          >
            {incident.ai_severity}
          </Typography>
        </Box>
        <Box mb={2}>
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>AI Category</strong>
          </Typography>
          <Typography
            data-testid="incident-ai-category"
            variant="body1"
            paragraph
          >
            {incident.ai_category}
          </Typography>
        </Box>
        <Box>
          <Typography variant="body2" color="textSecondary" gutterBottom>
            <strong>Created At</strong>
          </Typography>
          <Typography variant="body1" paragraph>
            {incident.created_at}
          </Typography>
        </Box>
      </Paper>
    </Container>
  );
}
