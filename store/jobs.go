package store

import (
	"database/sql"

	"github.com/leggettc18/job-tracker/types"
)

func (store *Storage) GetJobs() ([]types.Job, error) {
    rows, err := store.db.Query("SELECT * FROM jobs")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var jobs []types.Job
    for rows.Next() {
        job, err := scanJob(rows)
        if err != nil {
            return nil, err
        }
        jobs = append(jobs, job)
    }
    return jobs, nil
}

type JobRow struct {
    ID int
    Position string
    Company string
    Listing string
    CreatedAt string
    Status string
    InterviewDate sql.NullString
    AcceptedDate sql.NullString
}

func scanJob(row *sql.Rows) (types.Job, error) {
    var jobRow JobRow
    err := row.Scan(&jobRow.ID, &jobRow.Position, &jobRow.Company, &jobRow.Listing, &jobRow.CreatedAt, &jobRow.Status, &jobRow.InterviewDate, &jobRow.AcceptedDate)
    interviewDate := ""
    if jobRow.InterviewDate.Valid {
        interviewDate = jobRow.InterviewDate.String
    }
    acceptedDate := ""
    if jobRow.AcceptedDate.Valid {
        acceptedDate = jobRow.AcceptedDate.String
    }
    job := types.Job{
        ID: jobRow.ID,
        Position: jobRow.Position,
        Company: jobRow.Company,
        Listing: jobRow.Listing,
        CreatedAt: jobRow.CreatedAt,
        Status: jobRow.Status,
        InterviewDate: interviewDate,
        AcceptedDate: acceptedDate,
    }
    return job, err
}
