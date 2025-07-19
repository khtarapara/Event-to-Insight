import { useForm } from "react-hook-form";
import { createIncident } from "../../api/incident";
import { TextField, Button, Box, CircularProgress, Alert } from "@mui/material";
import { useState } from "react";

export default function IncidentForm({ onIncidentCreated, onCancel }) {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const onSubmit = async (data) => {
    setLoading(true);
    setError(null);
    try {
      await createIncident(data);
      reset();
      await onIncidentCreated(); // simple refresh
    } catch (err) {
      console.error("Error creating incident:", err);
      setError("Failed to create incident.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit(onSubmit)}
      noValidate
      sx={{ display: "flex", flexDirection: "column", gap: 2 }}
    >
      {error && <Alert severity="error">{error}</Alert>}
      <TextField
        label="Title"
        {...register("title", { required: true })}
        error={!!errors.title}
        helperText={errors.title && "Required"}
      />
      <TextField
        label="Description"
        multiline
        rows={3}
        {...register("description", { required: true })}
        error={!!errors.description}
        helperText={errors.description && "Required"}
      />
      <TextField
        label="Affected Service"
        {...register("affected_service", { required: true })}
        error={!!errors.affected_service}
        helperText={errors.affected_service && "Required"}
      />
      <Box sx={{ display: "flex", gap: 1 }}>
        <Button type="submit" variant="contained" disabled={loading}>
          {loading ? <CircularProgress size={24} /> : "Submit Incident"}
        </Button>
        <Button variant="outlined" onClick={() => reset()}>
          Reset
        </Button>
        <Button variant="text" onClick={() => onCancel()}>
          Cancel
        </Button>
      </Box>
    </Box>
  );
}
