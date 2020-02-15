package tests

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/notes"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/storage/tests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNote(t *testing.T) {
	// Mock Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Expected
	expectedContext := context.Background()
	expectedID := "my-id"
	expectedAuthor := "wow"
	expectedContent := "many"
	expectedTitle := "content"

	inputNoteItem := &storage.NoteItem{AuthorID: expectedAuthor, Content: expectedContent, Title: expectedTitle}
	expectedNoteItem := *inputNoteItem
	expectedNoteItem.ID = expectedID

	// Mocks
	mockStorage := tests.NewMockNoteStorage(ctrl)
	mockStorage.EXPECT().CreateNote(
		expectedContext,
		inputNoteItem,
	).Times(1).Return(&expectedNoteItem, nil)

	// Create Service
	service := notes.NewNoteService(mockStorage)

	// Create Request
	req := &notespb.CreateNoteRequest{
		Note: &notespb.Note{
			AuthorId: expectedAuthor,
			Content:  expectedContent,
			Title:    expectedTitle,
		},
	}

	// Act
	result, err := service.CreateNote(expectedContext, req)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, result.Note.Title, expectedTitle)
	assert.Equal(t, result.Note.AuthorId, expectedAuthor)
	assert.Equal(t, result.Note.Id, expectedID)
	assert.Equal(t, result.Note.Title, expectedTitle)
}

func TestReadNote(t *testing.T) {
	// Mock Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Expected
	expectedID := "my-id"
	expectedAuthor := "wow"
	expectedContent := "many"
	expectedTitle := "content"
	expectedNoteItem := &storage.NoteItem{ID: expectedID, AuthorID: expectedAuthor, Content: expectedContent, Title: expectedTitle}

	// Mocks
	mockStorage := tests.NewMockNoteStorage(ctrl)
	mockStorage.EXPECT().GetNote(
		gomock.Eq(context.Background()),
		gomock.Eq(expectedID),
	).Times(1).Return(expectedNoteItem, nil)

	// Create Service
	service := notes.NewNoteService(mockStorage)

	// Create Request
	req := &notespb.ReadNoteRequest{Id: expectedID}

	// Act
	result, err := service.ReadNote(context.Background(), req)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, result.Note.Title, expectedTitle)
	assert.Equal(t, result.Note.AuthorId, expectedAuthor)
	assert.Equal(t, result.Note.Id, expectedID)
	assert.Equal(t, result.Note.Title, expectedTitle)
}
