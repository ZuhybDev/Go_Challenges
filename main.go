package main

import "fmt"

func formatAppointment(patientName string, appointmentTime string) string {

	return fmt.Sprintf("Patient: %s - Appointment: %s\n", patientName, appointmentTime)
}

func main() {

	formatAppointment("", "2:30 PM")

}
