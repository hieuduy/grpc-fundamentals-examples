package com.elsevier.calculator;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.util.Arrays;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public class CalculatorClient {
  private static final String HOST = "localhost";
  private static final int PORT = 50051;

  private ManagedChannel channel;

  public static void main(String[] args) {
    CalculatorClient calculatorClient = new CalculatorClient();
    calculatorClient.run();
  }

  private void run() {
    channel = ManagedChannelBuilder.forAddress(HOST, PORT).usePlaintext().build();

    calculateSum(10, 20);
    decomposePrimeFactor(7192285800L);
    calculateSumEvenNumber(Arrays.asList(2, 3, 4, 9, 10, 18, 23, 30));
    findMaxNumber(Arrays.asList(4, 8, 2, 20, 12, 6, 32));

    channel.shutdown();
  }

  private void calculateSum(int firstNumber, int secondNumber) {
    CalculatorServiceGrpc.CalculatorServiceBlockingStub stub =
        CalculatorServiceGrpc.newBlockingStub(channel);

    Calculator.CalculateSumRequest request =
        Calculator.CalculateSumRequest.newBuilder()
            .setFirstNumber(firstNumber)
            .setSecondNumber(secondNumber)
            .build();

    Calculator.CalculateSumResponse response = stub.calculateSum(request);
    System.out.println("Response from CalculateSum: " + response.getSum());
  }

  private void decomposePrimeFactor(long number) {
    CalculatorServiceGrpc.CalculatorServiceBlockingStub stub =
        CalculatorServiceGrpc.newBlockingStub(channel);

    System.out.println("Response from DecomposePrimeFactor:");
    stub.decomposePrimeFactor(
            Calculator.DecomposePrimeFactorRequest.newBuilder().setNumber(number).build())
        .forEachRemaining(response -> System.out.println(response.getPrimeFactor()));
  }

  private void calculateSumEvenNumber(List<Integer> numbers) {
    CalculatorServiceGrpc.CalculatorServiceStub stub = CalculatorServiceGrpc.newStub(channel);
    CountDownLatch countDownLatch = new CountDownLatch(1);

    StreamObserver<Calculator.CalculateSumEvenNumberRequest> requestObserver =
        stub.calculateSumEvenNumber(
            new StreamObserver<Calculator.CalculateSumEvenNumberResponse>() {
              @Override
              public void onNext(Calculator.CalculateSumEvenNumberResponse response) {
                System.out.println("Response from CalculateSumEvenNumber: " + response.getSum());
              }

              @Override
              public void onError(Throwable t) {
                countDownLatch.countDown();
              }

              @Override
              public void onCompleted() {
                countDownLatch.countDown();
              }
            });

    numbers.forEach(
        number ->
            requestObserver.onNext(
                Calculator.CalculateSumEvenNumberRequest.newBuilder().setNumber(number).build()));
    requestObserver.onCompleted();

    try {
      countDownLatch.await(5, TimeUnit.SECONDS);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
  }

  private void findMaxNumber(List<Integer> numbers) {
    CalculatorServiceGrpc.CalculatorServiceStub stub = CalculatorServiceGrpc.newStub(channel);
    CountDownLatch countDownLatch = new CountDownLatch(1);

    StreamObserver<Calculator.FindMaxNumberRequest> requestObserver =
        stub.findMaxNumber(
            new StreamObserver<Calculator.FindMaxNumberResponse>() {
              @Override
              public void onNext(Calculator.FindMaxNumberResponse response) {
                System.out.println("New max number: " + response.getMaxNumber());
              }

              @Override
              public void onError(Throwable t) {
                countDownLatch.countDown();
              }

              @Override
              public void onCompleted() {
                countDownLatch.countDown();
              }
            });

    System.out.println("Response from FindMaxNumber:");
    numbers.forEach(
        number -> {
          requestObserver.onNext(
              Calculator.FindMaxNumberRequest.newBuilder().setNumber(number).build());
          try {
            Thread.sleep(100);
          } catch (InterruptedException e) {
            e.printStackTrace();
          }
        });
    requestObserver.onCompleted();

    try {
      countDownLatch.await(5, TimeUnit.SECONDS);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
  }
}
