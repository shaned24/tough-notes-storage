package notes

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/shaned24/tough-notes-storage/notes/notespb"
	"github.com/shaned24/tough-notes-storage/notes/server/storage"
	"github.com/shaned24/tough-notes-storage/notes/server/storage/tests"
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

	inputNoteItem := &storage.NoteItem{
		AuthorID: expectedAuthor,
		Content:  expectedContent,
		Title:    expectedTitle,
	}

	outputNoteItem := &storage.NoteItem{
		ID:       expectedID,
		AuthorID: expectedAuthor,
		Content:  expectedContent,
		Title:    expectedTitle,
	}

	// Mocks
	mockStorage := tests.NewMockNoteStorage(ctrl)
	mockStorage.EXPECT().CreateNote(expectedContext, inputNoteItem).Times(1).Return(outputNoteItem, nil)

	// Create Service
	service := NewNoteService(mockStorage)

	// Create Request
	req := &notespb.CreateNoteRequest{
		Note: &notespb.Note{
			AuthorId: expectedAuthor,
			Content:  expectedContent,
			Title:    expectedTitle,
		},
	}

	// Act
	result, _ := service.CreateNote(expectedContext, req)

	// Assert
	assert.Equal(t, result.Note.Title, expectedTitle)
	assert.Equal(t, result.Note.AuthorId, expectedAuthor)
	assert.Equal(t, result.Note.Id, expectedID)
	assert.Equal(t, result.Note.Title, expectedTitle)
}
