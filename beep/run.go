package beep

// AddNote adds notes
func (r *Run) AddNote(note Note) {
	r.Notes = append(r.Notes, note)
}

// AddNotes adds notes
func (r *Run) AddNotes(note ...Note) {
	r.Notes = append(r.Notes, note...)
}
