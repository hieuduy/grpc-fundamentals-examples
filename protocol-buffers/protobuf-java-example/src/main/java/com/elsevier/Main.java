package com.elsevier;

import com.elsevier.book.Book.BookMessage;

import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class Main {
  private static final String FILE_NAME = "book_message.bin";

  public static void main(String[] args) {
    BookMessage.Builder builder = BookMessage.newBuilder();
    BookMessage bookMessage =
        builder
            .setId(1)
            .setTitle("Java 8")
            .setAuthor("Oracle")
            .addTags("Development")
            .addTags("Java")
            .build();

    System.out.println(bookMessage);
    writeToFile(bookMessage, FILE_NAME);

    BookMessage bookMessageFromFile = readFromFile(FILE_NAME);
    System.out.println(bookMessageFromFile);
  }

  private static void writeToFile(BookMessage bookMessage, String fileName) {
    try {
      FileOutputStream outputStream = new FileOutputStream(fileName);
      bookMessage.writeTo(outputStream);
      outputStream.close();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }

  private static BookMessage readFromFile(String fileName) {
    try {
      FileInputStream fileInputStream = new FileInputStream(fileName);
      return BookMessage.parseFrom(fileInputStream);
    } catch (IOException e) {
      e.printStackTrace();
    }
    return null;
  }
}
